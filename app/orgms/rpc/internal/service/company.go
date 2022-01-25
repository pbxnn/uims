package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/orgms/rpc/internal/biz"
	"uims/app/orgms/rpc/internal/data/dao"

	pb "uims/api/orgms/rpc"
)

type CompanyService struct {
	pb.UnimplementedCompanyServer

	companyDO *biz.CompanyDO
	log       *log.Helper
}

func NewCompanyService(companyDO *biz.CompanyDO, logger log.Logger) *CompanyService {
	return &CompanyService{
		companyDO: companyDO,
		log:       log.NewHelper(log.With(logger, "service/company")),
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *pb.CreateCompanyReq) (*pb.CreateCompanyReply, error) {
	ext := "[]"
	if len(req.Ext) > 0 {
		extByte, err := json.Marshal(req.Ext)
		if err != nil {
			return nil, err
		}
		ext = string(extByte)
	}

	data := dao.OrgmsCompany{
		CompanyName: req.CompanyName,
		CompanyCode: req.CompanyCode,
		CompanyDesc: req.CompanyDesc,
		CompanyType: req.CompanyType,
		IsVirtual:   req.IsVirtual,
		DelFlag:     0,
		Ext:         ext,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}

	err := s.companyDO.Create(ctx, &data)
	if err != nil {
		return nil, err
	}

	reply := &pb.CreateCompanyReply{
		Data: &pb.CompanyInfo{
			CompanyId:   data.CompanyId,
			CompanyName: data.CompanyName,
			CompanyCode: data.CompanyCode,
			CompanyDesc: data.CompanyDesc,
			CompanyType: data.CompanyType,
			IsVirtual:   data.IsVirtual,
			StartTime:   data.StartTime,
			EndTime:     data.EndTime,
			Seq:         data.Seq,
			DelFlag:     data.DelFlag,
			CreatedAt:   data.CreatedAt.Unix(),
			UpdatedAt:   data.UpdatedAt.Unix(),
		},
	}

	json.Unmarshal([]byte(data.Ext), &reply.Data.Ext)
	return reply, err
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
