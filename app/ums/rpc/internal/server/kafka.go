package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/ums/rpc/internal/conf"
	"uims/app/ums/rpc/internal/service"
	"uims/pkg/kafka"
)

func NewKafkaConsumer(c *conf.KafkaConsumer, logger log.Logger) *kafka.KafkaSubClient {
	client := kafka.NewKafkaSub(kafka.KafkaConsumeConfig{
		Version: c.Version,
		Balance: c.RebalanceStrategy,
		Brokers: c.Brokers,
	}, logger)

	client.AddSubFunction(c.Topics, c.GroupId, service.UserActionHandler, nil)

	return client
}

func NewKafkaProducer(conf *conf.KafkaProducer) *kafka.KafkaPubClient {
	return kafka.NewKafkaPub(kafka.KafkaProducerConfig{
		Addr:    conf.Addrs,
		Version: conf.Version,
	})
}
