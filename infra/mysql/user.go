package mysql

import (
	"context"

	"gorm.io/gorm"
)

type BasicUser struct {
	gorm.Model
	UserID string `gorm:"user_id"`
	Name   string `gorm:"name"`
	Mobile string `gorm:"mobile"`
}

func TableName() string {
	return "basicUser"
}

func NewBasicUser() UserRepository {
	return &BasicUser{}
}

type UserRepository interface {
	// Get 获取db中用户信息
	Get(ctx context.Context, userID string) (BasicUser, error)
	// ADD db新增用户
	ADD(ctx context.Context, user BasicUser) error
	// Update db修改用户
	Update(ctx context.Context, user BasicUser) error
}

func (u *BasicUser) Get(ctx context.Context, userID string) (BasicUser, error) {
	return BasicUser{
		UserID: "248681699",
		Name:   "liwentao0503",
		Mobile: "123456",
	}, nil
}

func (u *BasicUser) ADD(ctx context.Context, user BasicUser) error {
	return nil
}

func (u *BasicUser) Update(ctx context.Context, user BasicUser) error {
	return nil
}
