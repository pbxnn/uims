package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
	"uims/app/ums/service/internal/conf"
	"uims/app/ums/service/internal/service"
)

type Message struct {
	Data        json.RawMessage   `json:"data"`
	SpanContext trace.SpanContext `json:"spanContext"`
}

type ConsumerHandler func(ctx context.Context, msg []byte) error

func NewKafkaConsumer(c *conf.Server, logger log.Logger, consumer sarama.Consumer, s *service.UserService) {

	addConsumer(consumer, 1, "uims_ums_create_user", func(ctx context.Context, msg []byte) error { return nil })
}

func addConsumer(c sarama.Consumer, partition int32, topic string, handler ConsumerHandler) {
	go func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		pc, err := c.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			//d.log.Warnf("init partition consumer err:%s", err.Error())
			fmt.Println(err)
			return
		}
		for {
			input := <-pc.Messages()
			msg := Message{}
			err := json.Unmarshal(input.Value, &msg)
			if err != nil {
				fmt.Println(err)
			}
			tracer := otel.Tracer("kratos")

			ctx := trace.ContextWithRemoteSpanContext(context.Background(), msg.SpanContext)
			_, span := tracer.Start(ctx, input.Topic, trace.WithSpanKind(trace.SpanKindConsumer), trace.WithTimestamp(time.Now()))

			err = handler(ctx, msg.Data)
			attrs := []attribute.KeyValue{
				attribute.String("topic", input.Topic),
				attribute.String("key", string(input.Key)),
				attribute.String("value", string(msg.Data)),
				attribute.Int64("offset", input.Offset),
				attribute.Int64("partition", int64(partition)),
			}

			if err != nil {
				attrs = append(attrs, attribute.String("error", err.Error()))
			}
			span.SetAttributes(attrs...)
			span.End()
		}

	}()
}
