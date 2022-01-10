package data

import (
	"context"
	"encoding/json"

	umsService "uims/api/ums/service"
	"uims/app/ums/service/internal/biz"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*umsService.GetUserReply, error) {
	return nil, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*umsService.GetUserByUsernameReply, error) {
	return nil, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*umsService.CreateUserReply, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	r.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: "uims_ums",
		Value: sarama.ByteEncoder(b),
	}
	return nil, nil

}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*umsService.SaveUserReply, error) {
	return nil, nil
}

func (r *userRepo) ListUser(ctx context.Context, u *biz.User) (*umsService.ListUserReply, error) {
	return nil, nil
}
