package data

import (
	"context"
	"encoding/json"

	"uims/app/ums/service/internal/biz"
	"uims/app/ums/service/internal/data/dao"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	um   *dao.UmsUserModel
	log  *log.Helper
}

func NewUserRepo(data *Data, um *dao.UmsUserModel, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		um:   um,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*dao.UmsUser, error) {
	userInfo, err := r.um.FindOne(ctx, id)
	r.log.WithContext(ctx).Infof("userInfo:%+v, err:%v", userInfo, err)
	return nil, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*dao.UmsUser, error) {
	return nil, nil
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*dao.UmsUser, error) {
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

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*dao.UmsUser, error) {
	return nil, nil
}

func (r *userRepo) ListUser(ctx context.Context, u *biz.User) ([]*dao.UmsUser, error) {
	return nil, nil
}
