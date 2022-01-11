package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	UmsUserModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	UmsUser struct {
		Id        int64     `gorm:"column:id"`
		Username  string    `gorm:"column:username"`
		DelFlag   int64     `gorm:"column:del_flag"`
		Status    int64     `gorm:"column:status"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}
)

func NewUmsUserModel(conn *gorm.DB, logger log.Logger) *UmsUserModel {
	return &UmsUserModel{
		conn:  conn,
		table: "ums_user",
		log:   log.NewHelper(log.With(logger, "module", "UmsUserModel")),
	}
}

func (m *UmsUserModel) Insert(ctx context.Context, data *UmsUser) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *UmsUserModel) BatchInsert(ctx context.Context, data []*UmsUser) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *UmsUserModel) FindOne(ctx context.Context, id int64) (*UmsUser, error) {
	data := &UmsUser{}
	err := m.conn.WithContext(ctx).Where("`id` = ?", id).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *UmsUserModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*UmsUser, error) {
	var data []*UmsUser
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *UmsUserModel) Update(ctx context.Context, data *UmsUser, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
