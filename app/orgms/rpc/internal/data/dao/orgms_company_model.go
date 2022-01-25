package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	OrgmsCompanyModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	OrgmsCompany struct {
		CompanyId   int64     `gorm:"primaryKey;column:company_id"` // 自增主键，单位id
		CompanyName string    `gorm:"column:company_name"`          // 单位名称
		CompanyCode string    `gorm:"column:company_code"`          // 单位代码
		CompanyDesc string    `gorm:"column:company_desc"`          // 单位描述
		CompanyType int64     `gorm:"column:company_type"`          // 单位类型
		Seq         int64     `gorm:"column:seq"`                   // 单位排序
		IsVirtual   int64     `gorm:"column:is_virtual"`            // 是否虚拟单位
		DelFlag     int64     `gorm:"column:del_flag"`              // 逻辑删除标识
		Ext         string    `gorm:"column:ext"`                   // 扩展信息
		StartTime   int64     `gorm:"column:start_time"`            // 单位开始时间
		EndTime     int64     `gorm:"column:end_time"`              // 单位结束时间
		CreatedAt   time.Time `gorm:"column:created_at"`            // 创建时间
		UpdatedAt   time.Time `gorm:"column:updated_at"`            // 更新时间
	}
)

func NewOrgmsCompanyModel(conn *gorm.DB, logger log.Logger) *OrgmsCompanyModel {
	return &OrgmsCompanyModel{
		conn:  conn,
		table: "orgms_company",
		log:   log.NewHelper(log.With(logger, "module", "OrgmsCompanyModel")),
	}
}

func (m *OrgmsCompanyModel) GetConn(ctx context.Context) *gorm.DB {
	return m.conn.WithContext(ctx)
}

func (m *OrgmsCompanyModel) Insert(ctx context.Context, data *OrgmsCompany) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *OrgmsCompanyModel) BatchInsert(ctx context.Context, data []*OrgmsCompany) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *OrgmsCompanyModel) FindOne(ctx context.Context, companyId int64) (*OrgmsCompany, error) {
	data := &OrgmsCompany{}
	err := m.conn.WithContext(ctx).Where("`company_id` = ?", companyId).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsCompanyModel) FindOneByCompanyName(ctx context.Context, companyName string) (*OrgmsCompany, error) {
	data := &OrgmsCompany{}
	err := m.conn.WithContext(ctx).Where("`company_name` = ?", companyName).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsCompanyModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*OrgmsCompany, error) {
	var data []*OrgmsCompany
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsCompanyModel) Update(ctx context.Context, data *OrgmsCompany, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}

func (m *OrgmsCompanyModel) UpdateMap(ctx context.Context, data map[string]interface{}, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
