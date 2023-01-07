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
	Get(ctx context.Context, userID string) (BasicUser, error)
	ADD(ctx context.Context, user BasicUser) error
	Update(ctx context.Context, user BasicUser) error
}

func (u BasicUser) Get(ctx context.Context, userID string) (BasicUser, error) {
	return BasicUser{
		UserID: "248681699",
		Name:   "wentao.li",
	}, nil
}

func (u BasicUser) ADD(ctx context.Context, user BasicUser) error {
	return nil
}

func (u BasicUser) Update(ctx context.Context, user BasicUser) error {
	return nil
}
