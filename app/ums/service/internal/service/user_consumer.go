package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	umsService "uims/api/ums/service"
	"uims/app/ums/service/internal/biz"
)

type UserConsumerService struct {
	umsService.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserConsumerService(uc *biz.UserUseCase, logger log.Logger) *UserConsumerService {
	return &UserConsumerService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "UserConsumerService"))}
}

func (s *UserConsumerService) UserCreatedConsumer(ctx context.Context, msg []byte) error {
	s.log.Infof("UserCreatedConsumer, msg:%s", string(msg))

	return nil
}
