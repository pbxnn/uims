package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	OrgmsUserExtraModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	OrgmsUserExtra struct {
		Id        int64     `gorm:"primaryKey;column:id"` // 自增主键
		Uid       int64     `gorm:"column:uid"`           // uid
		IsVirtual int64     `gorm:"column:is_virtual"`    // 是否虚拟用户
		AreaId    int64     `gorm:"column:area_id"`       // 地区id
		Ext       string    `gorm:"column:ext"`           // 扩展信息
		DelFlag   int64     `gorm:"column:del_flag"`      // 逻辑删除标识
		CreatedAt time.Time `gorm:"column:created_at"`    // 创建时间
		UpdatedAt time.Time `gorm:"column:updated_at"`    // 更新时间
	}
)

func NewOrgmsUserExtraModel(conn *gorm.DB, logger log.Logger) *OrgmsUserExtraModel {
	return &OrgmsUserExtraModel{
		conn:  conn,
		table: "orgms_user_extra",
		log:   log.NewHelper(log.With(logger, "module", "OrgmsUserExtraModel")),
	}
}

func (m *OrgmsUserExtraModel) Insert(ctx context.Context, data *OrgmsUserExtra) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *OrgmsUserExtraModel) BatchInsert(ctx context.Context, data []*OrgmsUserExtra) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *OrgmsUserExtraModel) FindOne(ctx context.Context, id int64) (*OrgmsUserExtra, error) {
	data := &OrgmsUserExtra{}
	err := m.conn.WithContext(ctx).Where("`id` = ?", id).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsUserExtraModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*OrgmsUserExtra, error) {
	var data []*OrgmsUserExtra
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsUserExtraModel) Update(ctx context.Context, data *OrgmsUserExtra, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
