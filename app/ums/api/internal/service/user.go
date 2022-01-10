package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/ums/api/internal/biz"

	pb "uims/api/ums/api"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	res, err := s.uc.GetUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	ret := &pb.GetUserReply{
		UserInfo: &pb.UserInfo{
			Uid:       res.UserInfo.Uid,
			Username:  res.UserInfo.Username,
			DelFlag:   res.UserInfo.DelFlag,
			Status:    res.UserInfo.Status,
			CreatedAt: res.UserInfo.CreatedAt,
			UpdatedAt: res.UserInfo.UpdatedAt,
		}}
	return ret, err
}
func (s *UserService) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameReq) (*pb.GetUserByUsernameReply, error) {
	return nil, nil
	//return s.uc.GetUser(ctx, req.Username)
}
func (s *UserService) Save(ctx context.Context, req *pb.SaveUserReq) (*pb.SaveUserReply, error) {
	return &pb.SaveUserReply{}, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	ret, err := s.uc.CreateUser(ctx, &biz.User{Username: req.Username})

	return &pb.CreateUserReply{
		Id:       ret.Id,
		Username: ret.Username,
	}, err
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
