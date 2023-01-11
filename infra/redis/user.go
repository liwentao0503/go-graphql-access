package redis

import "context"

// favorite classification 用户喜爱分类列表
type FCUser struct {
	FCIDs []uint64
}

// favorite classification 用户喜爱分类列表
var FCIDs []uint64

// KeyName hash表key名
func KeyName() string {
	return "favorite_classific:%s"
}

// NewFCUser 初始化redis 用户喜爱分类ID存储实体
func NewFCUser() UserFCRepository {
	return &FCUser{}
}

// UserFCRepository user喜爱分类列表存储防腐层
type UserFCRepository interface {
	// Get 获取db中用户信息
	Get(ctx context.Context, userID string) (FCUser, error)
	// Set db新增用户
	Set(ctx context.Context, IDs []uint64) error
}

func (u *FCUser) Get(ctx context.Context, userID string) (FCUser, error) {
	return FCUser{
		[]uint64{1, 2, 3, 4, 5},
	}, nil
}

func (u *FCUser) Set(ctx context.Context, IDs []uint64) error {
	return nil
}
