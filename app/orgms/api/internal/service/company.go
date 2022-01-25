package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/orgms/api/internal/biz"

	pb "uims/api/orgms/api"
)

type CompanyService struct {
	pb.UnimplementedCompanyServer

	companyDO *biz.CompanyDO
	log       *log.Helper
}

func NewCompanyService(companyDO *biz.CompanyDO, logger log.Logger) *CompanyService {
	return &CompanyService{
		companyDO: companyDO,
		log:       log.NewHelper(log.With(logger, "module", "service/company")),
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *pb.CreateCompanyReq) (*pb.CreateCompanyReply, error) {
	return s.companyDO.CreateCompany(ctx, req)
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
