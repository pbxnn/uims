package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	OrgmsUserDepartmentModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	OrgmsUserDepartment struct {
		Id           int64     `gorm:"primaryKey;column:id"` // 自增主键id
		Uid          int64     `gorm:"column:uid"`           // 用户id
		DepartmentId int64     `gorm:"column:department_id"` // 部门id
		CompanyId    int64     `gorm:"column:company_id"`    // 单位id
		IsAdmin      int64     `gorm:"column:is_admin"`      // 是否部门管理者
		IsShow       int64     `gorm:"column:is_show"`       // 是否显示
		DelFlag      int64     `gorm:"column:del_flag"`      // 逻辑删除标识
		CreatedAt    time.Time `gorm:"column:created_at"`    // 创建时间
		UpdatedAt    time.Time `gorm:"column:updated_at"`    // 更新时间
	}
)

func NewOrgmsUserDepartmentModel(conn *gorm.DB, logger log.Logger) *OrgmsUserDepartmentModel {
	return &OrgmsUserDepartmentModel{
		conn:  conn,
		table: "orgms_user_department",
		log:   log.NewHelper(log.With(logger, "module", "OrgmsUserDepartmentModel")),
	}
}

func (m *OrgmsUserDepartmentModel) Insert(ctx context.Context, data *OrgmsUserDepartment) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *OrgmsUserDepartmentModel) BatchInsert(ctx context.Context, data []*OrgmsUserDepartment) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *OrgmsUserDepartmentModel) FindOne(ctx context.Context, id int64) (*OrgmsUserDepartment, error) {
	data := &OrgmsUserDepartment{}
	err := m.conn.WithContext(ctx).Where("`id` = ?", id).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsUserDepartmentModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*OrgmsUserDepartment, error) {
	var data []*OrgmsUserDepartment
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsUserDepartmentModel) Update(ctx context.Context, data *OrgmsUserDepartment, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
