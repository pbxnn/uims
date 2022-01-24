package service

import (
	"context"

	pb "uims/api/orgms/api"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {

	return &pb.CreateUserReply{}, nil
}
func (s *UserService) BatchCreateUser(ctx context.Context, req *pb.BatchCreateUserReq) (*pb.BatchCreateUserReply, error) {
	return &pb.BatchCreateUserReply{}, nil
}
func (s *UserService) DelUser(ctx context.Context, req *pb.DelUserReq) (*pb.DelUserReply, error) {
	return &pb.DelUserReply{}, nil
}
func (s *UserService) GetUserInfoReq(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
