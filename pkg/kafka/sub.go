package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"reflect"
	"runtime"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
)

const KafkaBodyKey string = "KafkaMsg"
const CustomCtxKey string = "CustomCtx"

type CustomCtx struct {
	Handle string
	Type   string
	Desc   string
	Start  time.Time
}

type KafkaConsumeConfig struct {
	Version string   `yaml:"version"`
	Brokers []string `yaml:"brokers"`
	Balance string   `yaml:"balance"`
}

type KafkaSubClient struct {
	Brokers  []string
	Version  sarama.KafkaVersion
	Strategy sarama.BalanceStrategy
	ctx      context.Context
}

type KafkaBody struct {
	Msg         []byte
	SpanContext trace.SpanContext
}

func NewKafkaSub(c KafkaConsumeConfig) *KafkaSubClient {
	v, err := sarama.ParseKafkaVersion(c.Version)
	if err != nil {
		panic("Error parsing Kafka version: " + err.Error())
	}

	var s sarama.BalanceStrategy
	switch c.Balance {
	case "sticky":
		s = sarama.BalanceStrategySticky
	case "roundrobin":
		s = sarama.BalanceStrategyRoundRobin
	case "range":
		s = sarama.BalanceStrategyRange
	default:
		panic(fmt.Sprintf("Unrecognized consumer group partition assignor: %s", c.Balance))
	}

	return &KafkaSubClient{
		Version:  v,
		Brokers:  c.Brokers,
		Strategy: s,
		ctx:      context.Background(),
	}
}

type kafkaHandler func(context.Context) error

type KafkaConsumerOption struct {
	ConsumerFromNewest bool
}

func (c *KafkaSubClient) AddSubFunction(topics []string, groupID string, handler kafkaHandler, opts *KafkaConsumerOption) {
	config := sarama.NewConfig()
	config.Version = c.Version
	config.Consumer.Group.Rebalance.Strategy = c.Strategy

	if opts == nil || opts.ConsumerFromNewest == true {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	//config.Consumer.Return.Errors = true
	consumerGroup, err := sarama.NewConsumerGroup(c.Brokers, groupID, config)
	if err != nil {
		panic("NewConsumerGroup error: " + err.Error())
	}

	consumerHandler := &KafkaConsumerGroup{
		handler: handler,
		Client:  c,
		Ready:   make(chan bool),
	}

	ctx, _ := context.WithCancel(context.Background())

	go func() {
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := consumerGroup.Consume(ctx, topics, consumerHandler); err != nil {
				//zlog.Warn(nil, "Error from consumer: ", err.Error())
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumerHandler.Ready = make(chan bool, 0)
		}
	}()

	// Await till the consumer has been set up
	<-consumerHandler.Ready
	return
}

type KafkaConsumerGroup struct {
	Ready   chan bool
	handler func(ctx context.Context) error
	Client  *KafkaSubClient
	log     *log.Helper
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *KafkaConsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(c.Ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *KafkaConsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *KafkaConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		err := c.HandleMessage(message)
		if err != nil {
			continue
		}
		session.MarkMessage(message, "")
	}

	return nil
}

func (c *KafkaConsumerGroup) HandleMessage(message *sarama.ConsumerMessage) error {
	handlerName := reflect.TypeOf(c.handler).Name()
	customCtx := CustomCtx{
		Handle: handlerName,
		Type:   "kafka",
		Desc:   message.Topic,
	}
	ctx := context.WithValue(c.Client.ctx, CustomCtxKey, customCtx)

	var body KafkaBody
	if err := json.Unmarshal(message.Value, &body); err == nil {
		// kafka 消息发送的时候默认包装了一层msg，这里做个兼容。
		if body.Msg == nil {
			if err := json.Unmarshal(message.Value, &body.Msg); err != nil {
				c.log.Warn(ctx, "got unexpected value")
			}
		}
	} else {
		body.Msg = message.Value
	}

	tracer := otel.Tracer("uims")
	ctx = trace.ContextWithRemoteSpanContext(ctx, body.SpanContext)
	_, span := tracer.Start(ctx, message.Topic, trace.WithSpanKind(trace.SpanKindConsumer), trace.WithTimestamp(time.Now()))
	attrs := []attribute.KeyValue{
		attribute.String("topic", message.Topic),
		attribute.String("key", string(message.Key)),
		attribute.String("value", string(message.Value)),
		attribute.Int64("offset", message.Offset),
		attribute.Int64("partition", int64(message.Partition)),
	}
	span.SetAttributes(attrs...)

	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			info, _ := json.Marshal(map[string]interface{}{
				"time":   time.Now().Format("2006-01-02 15:04:05"),
				"level":  "error",
				"module": "stack",
				//"requestId": zlog.GetRequestID(ctx),
				"handle": handlerName,
			})
			fmt.Printf("%s\n-------------------stack-start-------------------\n%+v\n-------------------stack-end-------------------\n", string(info), r)
		}

		span.End()
	}()

	ctx = context.WithValue(ctx, KafkaBodyKey, body.Msg)

	err := c.handler(ctx)
	if err != nil {
		attrs = append(attrs, attribute.String("error", err.Error()))
	}

	// info 日志里打印出partition
	//.AddNotice(ctx, "partition", message.Partition)

	//ctx.CustomContext.Error = err
	//ctx.CustomContext.EndTime = time.Now()
	//middleware.LoggerAfterRun(ctx)
	return err
}

func GetKafkaMsg(ctx context.Context) []byte {
	return []byte(fmt.Sprint(ctx.Value(KafkaBodyKey)))
}
