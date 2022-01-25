package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type (
	UmsUserPwdModel struct {
		conn  *gorm.DB
		table string
		log   *log.Helper
	}

	UmsUserPwd struct {
		Id        int64     `gorm:"column:id"`
		Uid       int64     `gorm:"column:uid"`
		Pwd       string    `gorm:"column:pwd"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}
)

func NewUmsUserPwdModel(conn *gorm.DB, logger log.Logger) *UmsUserPwdModel {
	return &UmsUserPwdModel{
		conn:  conn,
		table: "ums_user_pwd",
		log:   log.NewHelper(log.With(logger, "module", "UmsUserPwdModel")),
	}
}

func (m *UmsUserPwdModel) Insert(ctx context.Context, data *UmsUserPwd) error {
	return m.conn.WithContext(ctx).Create(data).Error
}

func (m *UmsUserPwdModel) BatchInsert(ctx context.Context, data []*UmsUserPwd) error {
	return m.conn.WithContext(ctx).Create(&data).Error
}

func (m *UmsUserPwdModel) FindOne(ctx context.Context, id int64) (*UmsUserPwd, error) {
	data := &UmsUserPwd{}
	err := m.conn.WithContext(ctx).Where("`id` = ?", id).First(data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *UmsUserPwdModel) FindListByConds(ctx context.Context, conds map[string]interface{}) ([]*UmsUserPwd, error) {
	var data []*UmsUserPwd
	err := m.conn.WithContext(ctx).Where(conds).Find(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return data, err
}

func (m *UmsUserPwdModel) Update(ctx context.Context, data *UmsUserPwd, conds map[string]interface{}) error {
	return m.conn.WithContext(ctx).Where(conds).Updates(data).Error
}
