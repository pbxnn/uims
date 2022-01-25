package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	umsService "uims/api/ums/rpc"
	"uims/app/ums/api/internal/biz"
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
	return r.data.usc.GetUser(ctx, &umsService.GetUserReq{Uid: id})
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*umsService.GetUserByUsernameReply, error) {
	return r.data.usc.GetUserByUsername(ctx, &umsService.GetUserByUsernameReq{Username: username})
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*umsService.CreateUserReply, error) {
	return r.data.usc.CreateUser(ctx, &umsService.CreateUserReq{})
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*umsService.SaveUserReply, error) {
	return r.data.usc.Save(ctx, &umsService.SaveUserReq{})
}

func (r *userRepo) ListUser(ctx context.Context, u *biz.User) (*umsService.ListUserReply, error) {
	return r.data.usc.ListUser(ctx, &umsService.ListUserReq{})
}
