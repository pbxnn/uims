package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	OrgmsExtAttrModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	OrgmsExtAttr struct {
		Id           int64     `gorm:"primaryKey;column:id"` // 自增id
		AttrLevel    int64     `gorm:"column:attr_level"`    // 属性级别：1company，2department
		AttrType     int64     `gorm:"column:attr_type"`     // 属性类型：1单行，2列表
		AttrEnName   string    `gorm:"column:attr_en_name"`  // 英文属性名
		AttrCnName   string    `gorm:"column:attr_cn_name"`  // 中文属性名
		DefaultVaule string    `gorm:"column:default_vaule"` // 属性默认值
		IsRequired   int64     `gorm:"column:is_required"`   // 是否必填
		DelFlag      int64     `gorm:"column:del_flag"`      // 逻辑删除标识
		Seq          int64     `gorm:"column:seq"`           // 排序
		CreatedAt    time.Time `gorm:"column:created_at"`    // 创建时间
		UpdatedAt    time.Time `gorm:"column:updated_at"`    // 更新时间
	}
)

func NewOrgmsExtAttrModel(conn *gorm.DB, logger log.Logger) *OrgmsExtAttrModel {
	return &OrgmsExtAttrModel{
		conn:  conn,
		table: "orgms_ext_attr",
		log:   log.NewHelper(log.With(logger, "module", "OrgmsExtAttrModel")),
	}
}

func (m *OrgmsExtAttrModel) Insert(ctx context.Context, data *OrgmsExtAttr) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *OrgmsExtAttrModel) BatchInsert(ctx context.Context, data []*OrgmsExtAttr) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *OrgmsExtAttrModel) FindOne(ctx context.Context, id int64) (*OrgmsExtAttr, error) {
	data := &OrgmsExtAttr{}
	err := m.conn.WithContext(ctx).Where("`id` = ?", id).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsExtAttrModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*OrgmsExtAttr, error) {
	var data []*OrgmsExtAttr
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *OrgmsExtAttrModel) Update(ctx context.Context, data *OrgmsExtAttr, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
