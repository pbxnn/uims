package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	OrgmsDepartmentModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	OrgmsDepartment struct {
		DepartmentId   int64     `gorm:"primaryKey;column:department_id"` // 自增主键，部门id
		DepartmentName string    `gorm:"column:department_name"`          // 部门名称
		DepartmentDesc string    `gorm:"column:department_desc"`          // 部门描述
		DepartmentType int64     `gorm:"column:department_type"`          // 部门类型
		IsVirtual      int64     `gorm:"column:is_virtual"`               // 是否虚拟部门
		ParentId       int64     `gorm:"column:parent_id"`                // 上级部门id
		CompanyId      int64     `gorm:"column:company_id"`               // 单位id
		Depth          int64     `gorm:"column:depth"`                    // 当前部门层级
		Seq            int64     `gorm:"column:seq"`                      // 单位排序
		DelFlag        int64     `gorm:"column:del_flag"`                 // 逻辑删除标识
		Ext            string    `gorm:"column:ext"`                      // 扩展信息
		StartTime      int64     `gorm:"column:start_time"`               // 单位开始时间
		EndTime        int64     `gorm:"column:end_time"`                 // 单位结束时间
		CreatedAt      time.Time `gorm:"column:created_at"`               // 创建时间
		UpdatedAt      time.Time `gorm:"column:updated_at"`               // 更新时间
	}
)

func NewOrgmsDepartmentModel(conn *gorm.DB, logger log.Logger) *OrgmsDepartmentModel {
	return &OrgmsDepartmentModel{
		conn:  conn,
		table: "orgms_department",
		log:   log.NewHelper(log.With(logger, "module", "OrgmsDepartmentModel")),
	}
}

func (m *OrgmsDepartmentModel) Insert(ctx context.Context, data *OrgmsDepartment) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *OrgmsDepartmentModel) BatchInsert(ctx context.Context, data []*OrgmsDepartment) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *OrgmsDepartmentModel) FindOne(ctx context.Context, departmentId int64) (*OrgmsDepartment, error) {
	data := &OrgmsDepartment{}
	err := m.conn.WithContext(ctx).Where("`department_id` = ?", departmentId).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsDepartmentModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*OrgmsDepartment, error) {
	var data []*OrgmsDepartment
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsDepartmentModel) Update(ctx context.Context, data *OrgmsDepartment, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
