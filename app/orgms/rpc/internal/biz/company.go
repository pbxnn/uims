package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/orgms/rpc/internal/data/dao"
)

type CompanyDO struct {
	repo CompanyRepo
	log  *log.Helper
}

type CompanyRepo interface {
	CreateCompany(ctx context.Context, data *dao.OrgmsCompany) error
	BatchCreateCompany(ctx context.Context, data []*dao.OrgmsCompany) error
	GetCompany(ctx context.Context, companyId int64) (*dao.OrgmsCompany, error)
	GetCompanyList(ctx context.Context, pageSize, pageNum int64) ([]*dao.OrgmsCompany, int64, error)
	UpdateCompany(ctx context.Context, data *dao.OrgmsCompany) error
	OrderCompany(ctx context.Context, companyIds []int64) error
	DeleteCompany(ctx context.Context, companyId int64) error
}

func NewCompanyDO(repo CompanyRepo, logger log.Logger) *CompanyDO {
	return &CompanyDO{repo: repo, log: log.NewHelper(log.With(logger, "module", "biz/company"))}
}

func (do *CompanyDO) Create(ctx context.Context, data *dao.OrgmsCompany) error {
	return do.repo.CreateCompany(ctx, data)
}

func (do *CompanyDO) BatchCreate(ctx context.Context, data []*dao.OrgmsCompany) error {
	return do.repo.BatchCreateCompany(ctx, data)
}

func (do *CompanyDO) Get(ctx context.Context, companyId int64) (*dao.OrgmsCompany, error) {
	return do.repo.GetCompany(ctx, companyId)
}

func (do *CompanyDO) GetList(ctx context.Context, pageSize, pageNum int64) ([]*dao.OrgmsCompany, int64, error) {
	return do.repo.GetCompanyList(ctx, pageSize, pageNum)
}

func (do *CompanyDO) Update(ctx context.Context, data *dao.OrgmsCompany) error {
	return do.repo.UpdateCompany(ctx, data)
}

func (do *CompanyDO) Order(ctx context.Context, companyIds []int64) error {
	return do.repo.OrderCompany(ctx, companyIds)
}

func (do *CompanyDO) Delete(ctx context.Context, companyId int64) error {
	return do.repo.DeleteCompany(ctx, companyId)
}
