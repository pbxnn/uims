package service

import (
	"context"

	umsService "uims/api/ums/service"
	"uims/app/ums/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	umsService.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "UserService"))}
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

	userInfo, err := s.uc.CreateUser(ctx, &biz.User{Username: req.Username})
	if err != nil {
		return nil, err
	}

	rep := &umsService.CreateUserReply{
		Id:       userInfo.Id,
		Username: userInfo.Username,
	}
	return rep, err
}
func (s *UserService) ListUser(ctx context.Context, req *umsService.ListUserReq) (*umsService.ListUserReply, error) {
	return &umsService.ListUserReply{}, nil
}
