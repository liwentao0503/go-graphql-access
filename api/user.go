package api

import (
	"go-graphql-access/common"
	"go-graphql-access/domain"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

// UserInputType 用户信息输入
var UserInputType = graphql.FieldConfigArgument{
	"userID": &graphql.ArgumentConfig{Type: graphql.String, Description: "用户ID"},
}

// UserType 用户信息输出
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "userType",
	Fields: graphql.Fields{
		"userID": &graphql.Field{Type: graphql.String, Description: "用户ID"},
		"mobile": &graphql.Field{Type: graphql.String, Description: "用户手机号"},
		"name":   &graphql.Field{Type: graphql.String, Description: "用户姓名"},
		"fcIDs":  &graphql.Field{Type: graphql.NewList(graphql.Int), Description: "喜爱分类ID列表"},
	},
})

// UserReq 接入层 用户请求信息实体
type UserReq struct {
	UserID string `json:"userID"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
}

// UserRsp 接入层 用户返回信息实体
type UserRsp struct {
	UserID string `json:"userID"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
	// FCIDs （favorite classification） 喜爱分类ID列表
	FCIDs []uint64 `json:"fcIDs"`
}

var getUserInfo = &graphql.Field{
	Description: "查询用户信息",
	Type:        UserType,
	Args:        UserInputType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userReq, err := BuildUser(p.Args)
		if err != nil || userReq.ReadParamCheck() {
			return nil, common.ParamsError
		}
		user, err := domain.NewUserService().GetUser(p.Context, userReq.UserID,
			domain.WithBasic(), domain.WithFC())
		if err != nil {
			return nil, common.FormatError(err)
		}
		userRsp := UserRsp{}
		userRsp.ToEntity(user)
		return userRsp, nil
	},
}

var addUserInfo = &graphql.Field{
	Description: "新增用户信息",
	Type:        UserType,
	Args:        UserInputType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userReq, err := BuildUser(p.Args)
		if err != nil || userReq.WriteParamCheck() {
			return nil, common.ParamsError
		}
		return nil, common.FormatError(domain.NewUserService().
			ADDBasicUser(p.Context, userReq.FromEntity()))
	},
}

var updateUserInfo = &graphql.Field{
	Description: "更新用户信息",
	Type:        UserType,
	Args:        UserInputType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		user, err := BuildUser(p.Args)
		if err != nil || user.WriteParamCheck() {
			return nil, common.ParamsError
		}
		return nil, common.FormatError(domain.NewUserService().
			UpdateBasicUser(p.Context, user.FromEntity()))
	},
}

// BuildUser 构建接入层User
func BuildUser(args map[string]interface{}) (*UserReq, error) {
	userReq := UserReq{}
	if err := mapstructure.Decode(args, &userReq); err != nil {
		return nil, err
	}
	return &userReq, nil
}

// ReadParamCheck true参数校验失败 false参数校验成功
func (u *UserReq) ReadParamCheck() bool {
	if u.UserID == "" || len(u.UserID) > 25 {
		return true
	}
	return false
}

// WriteParamCheck true参数校验失败 false参数校验成功
func (u *UserReq) WriteParamCheck() bool {
	if u.UserID == "" || u.Mobile == "" {
		return true
	}
	return false
}

// FromEntity 从接入层实体转化为业务实体
func (u *UserReq) FromEntity() domain.User {
	return domain.User{
		UserID: u.UserID,
		Mobile: u.Mobile,
		Name:   u.Name,
	}
}

// ToEntity 从业务实体转化为接入层实体
func (u *UserRsp) ToEntity(user domain.User) {
	u.UserID = user.UserID
	u.Mobile = user.Mobile
	u.Name = user.Name
	u.FCIDs = user.FCIDs
}
