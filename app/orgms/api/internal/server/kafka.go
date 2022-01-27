package server

import (
	"uims/app/orgms/api/internal/conf"
	"uims/pkg/kafka"
)

func NewKafkaProducer(conf *conf.KafkaProducer) *kafka.KafkaPubClient {
	return kafka.NewKafkaPub(kafka.KafkaProducerConfig{
		Addr:    conf.Addrs,
		Version: conf.Version,
	})
}
