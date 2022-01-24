package service

import (
	"context"

	pb "uims/api/orgms/api"
)

type CompanyService struct {
	pb.UnimplementedCompanyServer
}

func NewCompanyService() *CompanyService {
	return &CompanyService{}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *pb.CreateCompanyReq) (*pb.CreateCompanyReply, error) {
	return &pb.CreateCompanyReply{}, nil
}
func (s *CompanyService) BatchCreateCompany(ctx context.Context, req *pb.CreateCompanyReq) (*pb.BatchCreateCompanyReply, error) {
	return &pb.BatchCreateCompanyReply{}, nil
}
func (s *CompanyService) GetCompanyList(ctx context.Context, req *pb.GetCompanyListReq) (*pb.GetCompanyListReply, error) {
	return &pb.GetCompanyListReply{}, nil
}
func (s *CompanyService) GetCompany(ctx context.Context, req *pb.GetCompanyReq) (*pb.GetCompanyReply, error) {
	return &pb.GetCompanyReply{}, nil
}
func (s *CompanyService) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyReq) (*pb.UpdateCompanyReply, error) {
	return &pb.UpdateCompanyReply{}, nil
}
func (s *CompanyService) DeleteCompany(ctx context.Context, req *pb.DelCompanyReq) (*pb.DelCompanyReply, error) {
	return &pb.DelCompanyReply{}, nil
}
func (s *CompanyService) OrderCompany(ctx context.Context, req *pb.OrderCompanyReq) (*pb.OrderCompanyReply, error) {
	return &pb.OrderCompanyReply{}, nil
}
