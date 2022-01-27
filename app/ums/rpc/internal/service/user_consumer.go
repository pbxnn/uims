package service

import (
	"context"
	"encoding/json"
	"fmt"

	"uims/api/ums/rpc"
	"uims/pkg/kafka"
)

//
//type UserConsumerService struct {
//	umsService.UnimplementedUserServer
//
//	uc  *biz.UserUseCase
//	log *log.Helper
//}
//
//func NewUserConsumerService(uc *biz.UserUseCase, logger log.Logger) *UserConsumerService {
//	return &UserConsumerService{
//		uc:  uc,
//		log: log.NewHelper(log.With(logger, "module", "UserConsumerService"))}
//}
//
//func (s *UserConsumerService) UserCreatedConsumer(ctx context.Context, msg []byte) error {
//	s.log.Infof("UserCreatedConsumer, msg:%s", string(msg))
//
//	return nil
//}

func UserActionHandler(ctx context.Context) error {
	body := kafka.GetKafkaMsg(ctx)
	msg := &rpc.UserActionMsg{}
	if err := json.Unmarshal(body, msg); err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
