package domain

import (
	"context"

	"go-graphql-access/infra/mysql"
	"go-graphql-access/infra/redis"
)

var _ UserService = (*User)(nil)

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

type Option func(*Options)

// Options  每次请求的上下文数据
type Options struct {
	WithBasic  bool
	WithFC     bool
	WithFollow bool
	WithFans   bool
	WithPost   bool
}

// WithBasic 获取基本用户信息
func WithBasic() Option {
	return func(o *Options) {
		o.WithBasic = true
	}
}

// WithFC 获取用户喜爱分类ID列表
func WithFC() Option {
	return func(o *Options) {
		o.WithFC = true
	}
}

// WithFollow 获取用户关注列表
func WithFollow() Option {
	return func(o *Options) {
		o.WithFollow = true
	}
}

// WithFans 获取用户粉丝列表
func WithFans() Option {
	return func(o *Options) {
		o.WithFans = true
	}
}

// WithPost 获取用户发帖列表
func WithPost() Option {
	return func(o *Options) {
		o.WithFans = true
	}
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
	// GetUser 获取用户信息
	GetUser(ctx context.Context, userID string, opts ...Option) (User, error)
	// ADDBasicUser 新增用户基础信息
	ADDBasicUser(ctx context.Context, user User) error
	// UpdateBasicUser 修改用户实体信息
	UpdateBasicUser(ctx context.Context, user User) error
	// SetFCUser 设置喜爱分类ID
	SetFCUser(ctx context.Context, IDs []uint64) error
}

// NewUserService 初始化用户服务变量
func NewUserService() UserService {
	return &User{}
}

func (u *User) GetUser(ctx context.Context, userID string, opts ...Option) (User, error) {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	if options.WithBasic {
		basicUser, err := mysql.NewBasicUser().Get(ctx, userID)
		if err != nil {
			return User{}, err
		}
		u.ToBasicUserEntity(basicUser)
	}

	if options.WithFC {
		fcUser, err := redis.NewFCUser().Get(ctx, userID)
		if err != nil {
			return User{}, err
		}
		u.ToFCUserEntity(fcUser)
	}
	return *u, nil
}

func (u *User) ADDBasicUser(ctx context.Context, user User) error {
	return mysql.NewBasicUser().ADD(ctx, user.FromBasicUserEntity())
}

func (u *User) UpdateBasicUser(ctx context.Context, user User) error {
	return mysql.NewBasicUser().Update(ctx, user.FromBasicUserEntity())
}

func (u *User) SetFCUser(ctx context.Context, IDs []uint64) error {
	return redis.NewFCUser().Set(ctx, IDs)
}
