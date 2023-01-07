package domain

import (
	"code_struct/infra/mysql"
	"context"
)

type User struct {
	UserID string `json:"userID"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
	// FCIDs （favorite classification） 喜爱分类ID列表
	FCIDs []uint64 `json:"fcIDs"`
	// FollowUIDs 关注用户列表
	FollowUIDs []string `json:"followUIDs"`
	// FansUIDs 粉丝列表 （最多取近1k）
	FansUIDs []string `json:"fansUIDs"`
	// PostIDS 发帖列表
	PostIDS []uint64 `json:"postIDs"`
}

// FromBasicEntity 从业务实体转化为db实体
func (u *User) FromBasicUserEntity() mysql.BasicUser {
	return mysql.BasicUser{
		UserID: u.UserID,
		Mobile: u.Mobile,
		Name:   u.Name,
	}
}

// ToBasicEntity 从db实体转化为业务实体
func (u *User) ToBasicUserEntity(user mysql.BasicUser) User {
	return User{
		UserID: user.UserID,
		Mobile: user.Mobile,
		Name:   user.Name,
	}
}

type UserService interface {
	GetBasicUser(ctx context.Context, userID string) (User, error)
	ADDBasicUser(ctx context.Context, user User) error
	UpdateBasicUser(ctx context.Context, user User) error
}

func NewUserService() UserService {
	return &User{}
}

func (u User) GetBasicUser(ctx context.Context, userID string) (User, error) {
	userInfo, err := mysql.NewBasicUser().Get(ctx, userID)
	if err != nil {
		return User{}, err
	}
	return u.ToBasicUserEntity(userInfo), nil
}

func (u User) ADDBasicUser(ctx context.Context, user User) error {
	return mysql.NewBasicUser().ADD(ctx, user.FromBasicUserEntity())
}

func (u User) UpdateBasicUser(ctx context.Context, user User) error {
	return mysql.NewBasicUser().Update(ctx, user.FromBasicUserEntity())
}
