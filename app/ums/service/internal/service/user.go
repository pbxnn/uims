package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/ums/service/internal/biz"

	umsService "uims/api/ums/service"
)

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/server-service"))}
}

func (s *UserService) GetUser(ctx context.Context, req *umsService.GetUserReq) (*umsService.GetUserReply, error) {
	return &umsService.GetUserReply{
		UserInfo: &umsService.UserInfo{Uid: req.Uid, Username: "zxn"},
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, req *umsService.GetUserByUsernameReq) (*umsService.GetUserByUsernameReply, error) {
	return &umsService.GetUserByUsernameReply{}, nil
}
func (s *UserService) Save(ctx context.Context, req *umsService.SaveUserReq) (*umsService.SaveUserReply, error) {
	return &umsService.SaveUserReply{}, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *umsService.CreateUserReq) (*umsService.CreateUserReply, error) {

	return s.uc.CreateUser(ctx, &biz.User{Username: req.Username})
}
func (s *UserService) ListUser(ctx context.Context, req *umsService.ListUserReq) (*umsService.ListUserReply, error) {
	return &umsService.ListUserReply{}, nil
}
