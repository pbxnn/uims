package service

import (
	"context"

	pb "uims/api/orgms/api"
)

type DepartmentService struct {
	pb.UnimplementedDepartmentServer
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{}
}

func (s *DepartmentService) GetDepartment(ctx context.Context, req *pb.GetDepartmentReq) (*pb.GetDepartmentReply, error) {
	return &pb.GetDepartmentReply{}, nil
}
func (s *DepartmentService) GetDepartmentList(ctx context.Context, req *pb.GetDepartmentListReq) (*pb.GetCompanyListReply, error) {
	return &pb.GetCompanyListReply{}, nil
}
func (s *DepartmentService) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentReq) (*pb.CreateDepartmentReply, error) {
	return &pb.CreateDepartmentReply{}, nil
}
func (s *DepartmentService) BatchCreateDepartment(ctx context.Context, req *pb.BatchCreateDepartmentReq) (*pb.BatchCreateDepartmentReply, error) {
	return &pb.BatchCreateDepartmentReply{}, nil
}
func (s *DepartmentService) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentReq) (*pb.UpdateDepartmentReply, error) {
	return &pb.UpdateDepartmentReply{}, nil
}
func (s *DepartmentService) DeleteDepartment(ctx context.Context, req *pb.DelDepartmentReq) (*pb.DelDepartmentReply, error) {
	return &pb.DelDepartmentReply{}, nil
}
func (s *DepartmentService) AssignDepartmentUser(ctx context.Context, req *pb.AssignDepartmentUserReq) (*pb.AssignDepartmentUserReply, error) {
	return &pb.AssignDepartmentUserReply{}, nil
}
func (s *DepartmentService) OrderDepartment(ctx context.Context, req *pb.OrderDepartmentReq) (*pb.OrderDepartmentReply, error) {
	return &pb.OrderDepartmentReply{}, nil
}
func (s *DepartmentService) MoveDepartment(ctx context.Context, req *pb.MoveDepartmentReq) (*pb.MoveDepartmentReply, error) {
	return &pb.MoveDepartmentReply{}, nil
}
func (s *DepartmentService) MergeDepartment(ctx context.Context, req *pb.MergeDepartmentReq) (*pb.MergeDepartmentReply, error) {
	return &pb.MergeDepartmentReply{}, nil
}
