package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type KafkaProducerConfig struct {
	Service string   `yaml:"service"`
	Addr    []string `yaml:"addr"`
	Version string   `yaml:"version"`

	SASL struct {
		Enable    bool   `yaml:"enable"`
		Handshake bool   `yaml:"handshake"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
	} `yaml:"sasl"`

	TLS struct {
		Enable                bool   `yaml:"enable"`
		CA                    string `yaml:"ca"`
		Cert                  string `yaml:"cert"`
		Key                   string `yaml:"key"`
		InsecureSkipTLSVerify bool   `yaml:"insecure_skip_tls_verify"`
	} `yaml:"tls"`
}
type KafkaPubClient struct {
	Conf     KafkaProducerConfig
	producer sarama.SyncProducer
}

const kafkaPrefix = "@@kafkapub."

func (conf *KafkaProducerConfig) GetKafkaConfig() (*sarama.Config, error) {
	//secret.CommonSecretChange(kafkaPrefix, *conf, conf)

	config := sarama.NewConfig()
	v, err := sarama.ParseKafkaVersion(conf.Version)
	if err != nil {
		return nil, err
	}
	config.Version = v
	if conf.SASL.Enable {
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = conf.SASL.Handshake
		config.Net.SASL.User = conf.SASL.User
		config.Net.SASL.Password = conf.SASL.Password
	}
	if conf.TLS.Enable {
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = &tls.Config{
			RootCAs:            x509.NewCertPool(),
			InsecureSkipVerify: conf.TLS.InsecureSkipTLSVerify,
		}
		if conf.TLS.CA != "" {
			ca, err := ioutil.ReadFile(conf.TLS.CA)
			if err != nil {
				panic("kafka pub CA error: %v" + err.Error())
			}
			config.Net.TLS.Config.RootCAs.AppendCertsFromPEM(ca)
		}
	}

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	return config, nil
}

func NewKafkaPub(conf KafkaProducerConfig) *KafkaPubClient {
	saramaConfig, err := conf.GetKafkaConfig()
	if err != nil {
		panic("kafka pub version error: %v" + err.Error())
	}

	producer, err := sarama.NewSyncProducer(conf.Addr, saramaConfig)
	if err != nil {
		panic("kafka pub new producer error: %v" + err.Error())
	}

	c := &KafkaPubClient{
		Conf:     conf,
		producer: producer,
	}
	return c
}

func (client *KafkaPubClient) CloseProducer() error {
	if client.producer != nil {
		return client.producer.Close()
	}
	return nil
}

func (client *KafkaPubClient) Pub(ctx context.Context, topic string, msg []byte) (int32, int64, error) {
	if client.producer == nil {
		return 0, 0, errors.New("kafka producer not init")
	}
	kafkaBody := KafkaBody{
		Msg:         msg,
		SpanContext: trace.SpanContextFromContext(ctx),
	}
	body, err := json.Marshal(kafkaBody)
	if err != nil {
		return 0, 0, err
	}

	//start := time.Now()
	kafkaMsg := &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(body)}
	partition, offset, err := client.producer.SendMessage(kafkaMsg)
	//end := time.Now()

	if _, ok := transport.FromServerContext(ctx); ok {
		tracer := otel.Tracer("uims")
		_, span := tracer.Start(ctx, kafkaMsg.Topic, trace.WithSpanKind(trace.SpanKindProducer), trace.WithTimestamp(time.Now()))
		defer span.End()

		msgKey, _ := kafkaMsg.Key.Encode()
		msgValue, _ := kafkaMsg.Value.Encode()
		attrs := []attribute.KeyValue{
			attribute.String("topic", kafkaMsg.Topic),
			attribute.String("key", string(msgKey)),
			attribute.String("value", string(msgValue)),
			attribute.Int64("offset", offset),
			attribute.Int64("partition", int64(partition)),
		}

		if err != nil {
			attrs = append(attrs, attribute.String("error", err.Error()))
		}
		span.SetAttributes(attrs...)
	}

	//infoMsg := "kafka pub success"
	//if err != nil {
	//	ralCode = -1
	//	infoMsg = err.Error()
	//	zlog.ErrorLogger(ctx, "kafka pub error: "+infoMsg, zlog.String(zlog.TopicType, zlog.LogNameModule))
	//}
	//
	//fields := []zlog.Field{
	//	zlog.String(zlog.TopicType, zlog.LogNameModule),
	//	zlog.String("requestId", zlog.GetRequestID(ctx)),
	//	zlog.String("localIp", env.LocalIP),
	//	zlog.String("remoteAddr", client.Conf.Addr),
	//	zlog.String("service", client.Conf.Service),
	//	zlog.Int32("partition", partition),
	//	zlog.Int64("offset", offset),
	//	zlog.Int("ralCode", ralCode),
	//	zlog.String("requestStartTime", utils.GetFormatRequestTime(start)),
	//	zlog.String("requestEndTime", utils.GetFormatRequestTime(end)),
	//	zlog.Float64("cost", utils.GetRequestCost(start, end)),
	//}
	//
	//zlog.InfoLogger(nil, infoMsg, fields...)

	return partition, offset, err
}
