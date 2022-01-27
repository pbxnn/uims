package data

import (
	"context"
	"encoding/json"

	"uims/app/ums/rpc/internal/biz"
	"uims/app/ums/rpc/internal/data/dao"

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
	msg, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	_, _, err = r.data.kp.Pub(ctx, "uims_ums_user_create", msg)

	return nil, err

}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*dao.UmsUser, error) {
	return nil, nil
}

func (r *userRepo) ListUser(ctx context.Context, u *biz.User) ([]*dao.UmsUser, error) {
	return nil, nil
}
