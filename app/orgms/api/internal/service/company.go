package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uims/api/ums/rpc"
	"uims/app/orgms/api/internal/biz"
	"uims/app/orgms/api/internal/pkg"
	"uims/pkg/kafka"

	pb "uims/api/orgms/api"
)

type CompanyService struct {
	pb.UnimplementedCompanyServer

	companyDO *biz.CompanyDO
	kp        *kafka.KafkaPubClient
	log       *log.Helper
}

func NewCompanyService(companyDO *biz.CompanyDO, kp *kafka.KafkaPubClient, logger log.Logger) *CompanyService {
	return &CompanyService{
		companyDO: companyDO,
		kp:        kp,
		log:       log.NewHelper(log.With(logger, "module", "service/company")),
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *pb.CreateCompanyReq) (*pb.CreateCompanyReply, error) {
	rep, err := s.companyDO.CreateCompany(ctx, req)

	pkg.NewUserAction(s.kp, 111, rpc.ACTION_TYPE_ORGMS_CREATE_COMPANY, map[string]interface{}{"req": req}).Send(ctx)
	return rep, err
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
