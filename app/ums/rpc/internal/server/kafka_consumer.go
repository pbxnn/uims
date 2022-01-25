package server

import (
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/ums/rpc/internal/conf"
	"uims/app/ums/rpc/internal/pkg/kafka"
	"uims/app/ums/rpc/internal/service"
)

func NewKafkaConsumer(c *conf.Server, logger log.Logger, consumer sarama.Consumer, s *service.UserConsumerService) {
	kafka.Consume(consumer, 1, "uims_ums_create_user", s.UserCreatedConsumer)
}
