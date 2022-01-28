package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"io/ioutil"
	"time"
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

	p, err := sarama.NewSyncProducer(conf.Addr, saramaConfig)
	if err != nil {
		panic("kafka pub new producer error: %v" + err.Error())
	}

	traceProvider := otel.GetTracerProvider()
	propagators := otel.GetTextMapPropagator()
	opts := []otelsarama.Option{
		otelsarama.WithTracerProvider(traceProvider),
		otelsarama.WithPropagators(propagators),
	}
	producer := otelsarama.WrapSyncProducer(saramaConfig, p, opts...)

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
		Msg: msg,
	}

	body, err := json.Marshal(kafkaBody)
	if err != nil {
		return 0, 0, err
	}

	kafkaMsg := &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(body)}
	propagators := propagation.TraceContext{}
	propagators.Inject(ctx, otelsarama.NewProducerMessageCarrier(kafkaMsg))

	partition, offset, err := client.producer.SendMessage(kafkaMsg)

	if _, ok := transport.FromServerContext(ctx); ok {
		tracer := otel.Tracer("uims")
		_, span := tracer.Start(ctx, "kafka_produce/"+kafkaMsg.Topic, trace.WithSpanKind(trace.SpanKindProducer), trace.WithTimestamp(time.Now()))
		defer span.End()

		attrs := []attribute.KeyValue{
			attribute.String("topic", kafkaMsg.Topic),
			attribute.String("msg", string(msg)),
			attribute.String("traceId", span.SpanContext().TraceID().String()),
			attribute.String("spanId", span.SpanContext().SpanID().String()),
			attribute.Int64("offset", offset),
			attribute.Int64("partition", int64(partition)),
		}

		if err != nil {
			attrs = append(attrs, attribute.String("error", err.Error()))
		}
		span.SetAttributes(attrs...)
	}

	return partition, offset, err
}
