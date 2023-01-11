package domain

import (
	"context"

	"go-graphql-access/infra/mysql"
	"go-graphql-access/infra/redis"
)

// User 用户相关信息实体类
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
func (u *User) ToBasicUserEntity(user mysql.BasicUser) {
	u.UserID = user.UserID
	u.Mobile = user.Mobile
	u.Name = user.Name
}

// FromBasicEntity 从业务实体转化为db实体
func (u *User) FromFCUserEntity() redis.FCUser {
	return redis.FCUser{
		FCIDs: u.FCIDs,
	}
}

// ToBasicEntity 从db实体转化为业务实体
func (u *User) ToFCUserEntity(user redis.FCUser) {
	u.FCIDs = user.FCIDs
}

// UserService 用户服务防腐层接口
type UserService interface {
	// GetBasicUser 获取用户基础信息
	GetBasicUser(ctx context.Context, userID string) (User, error)
	// ADDBasicUser 新增用户基础信息
	ADDBasicUser(ctx context.Context, user User) error
	// UpdateBasicUser 修改用户实体信息
	UpdateBasicUser(ctx context.Context, user User) error
	// GetFCUser 获取喜爱分类ID列表
	GetFCUser(ctx context.Context, userID string) (User, error)
	// SetFCUser 设置喜爱分类ID
	SetFCUser(ctx context.Context, IDs []uint64) error
}

// NewUserService 初始化用户服务变量
func NewUserService() UserService {
	return &User{}
}

func (u *User) GetBasicUser(ctx context.Context, userID string) (User, error) {
	basicUser, err := mysql.NewBasicUser().Get(ctx, userID)
	if err != nil {
		return User{}, err
	}
	u.ToBasicUserEntity(basicUser)
	return *u, nil
}

func (u *User) ADDBasicUser(ctx context.Context, user User) error {
	return mysql.NewBasicUser().ADD(ctx, user.FromBasicUserEntity())
}

func (u *User) UpdateBasicUser(ctx context.Context, user User) error {
	return mysql.NewBasicUser().Update(ctx, user.FromBasicUserEntity())
}

func (u *User) GetFCUser(ctx context.Context, userID string) (User, error) {
	fcUser, err := redis.NewFCUser().Get(ctx, userID)
	if err != nil {
		return User{}, err
	}
	u.ToFCUserEntity(fcUser)
	return *u, nil
}

func (u *User) SetFCUser(ctx context.Context, IDs []uint64) error {
	return redis.NewFCUser().Set(ctx, IDs)
}
