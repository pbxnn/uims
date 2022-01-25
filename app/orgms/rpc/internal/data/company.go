package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uims/app/orgms/rpc/internal/biz"
	"uims/app/orgms/rpc/internal/data/dao"
)

type companyRepo struct {
	data *Data
	cm   *dao.OrgmsCompanyModel
	log  *log.Helper
}

func NewCompanyRepo(data *Data, cm *dao.OrgmsCompanyModel, logger log.Logger) biz.CompanyRepo {
	return &companyRepo{
		data: data,
		cm:   cm,
		log:  log.NewHelper(log.With(logger, "module", "data/company")),
	}
}

func (repo *companyRepo) CreateCompany(ctx context.Context, data *dao.OrgmsCompany) error {
	err := repo.cm.Insert(ctx, data)
	return err
}

func (repo *companyRepo) BatchCreateCompany(ctx context.Context, data []*dao.OrgmsCompany) error {
	err := repo.cm.BatchInsert(ctx, data)
	return err
}

func (repo *companyRepo) GetCompany(ctx context.Context, companyId int64) (*dao.OrgmsCompany, error) {
	data, err := repo.cm.FindOne(ctx, companyId)
	return data, err
}

func (repo *companyRepo) GetCompanyList(ctx context.Context, pageSize, pageNum int64) ([]*dao.OrgmsCompany, int64, error) {
	var data []*dao.OrgmsCompany
	var total int64
	if err := repo.data.db.WithContext(ctx).Offset(int(pageSize * pageNum)).Limit(int(pageNum)).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	if err := repo.data.db.WithContext(ctx).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (repo *companyRepo) UpdateCompany(ctx context.Context, data *dao.OrgmsCompany) error {
	conds := map[string]interface{}{
		"company_id": data.CompanyId,
	}
	err := repo.cm.Update(ctx, data, conds)

	return err
}

func (repo *companyRepo) OrderCompany(ctx context.Context, companyIds []int64) error {
	tx := repo.cm.GetConn(ctx).Begin()
	if err := tx.Error; err != nil {
		return err
	}
	defer tx.Rollback()

	for seq, companyId := range companyIds {
		if err := tx.Where("company_id = ?", companyId).Update("seq", seq).Error; err != nil {
			return err
		}
	}

	err := tx.Commit().Error
	return err

}

func (repo *companyRepo) DeleteCompany(ctx context.Context, companyId int64) error {
	conds := map[string]interface{}{
		"company_id": companyId,
	}

	data := map[string]interface{}{
		"del_flag": 1,
	}
	err := repo.cm.UpdateMap(ctx, data, conds)
	return err
}
