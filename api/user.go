package api

import (
	"go-graphql-access/common"
	"go-graphql-access/domain"

	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

// UserInfoInput 用户信息输入
var UserInfoInput = graphql.FieldConfigArgument{
	"userID": &graphql.ArgumentConfig{Type: graphql.String, Description: "userID"},
}

// UserInfoType 用户信息输出
var UserInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "userInfoType",
	Fields: graphql.Fields{
		"userID": &graphql.Field{Type: graphql.String, Description: "用户ID"},
		"mobile": &graphql.Field{Type: graphql.String, Description: "用户手机号"},
		"name":   &graphql.Field{Type: graphql.String, Description: "用户姓名"},
	},
})

// User 接入层用户信息实体
type User struct {
	UserID string `json:"userID"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
}

var getUserInfo = &graphql.Field{
	Description: "查询用户信息",
	Type:        UserInfoType,
	Args:        UserInfoInput,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		user, err := BuildUser(p.Args)
		if err != nil || user.ReadParamCheck() {
			return nil, common.ParamsError
		}
		userInfo, err := domain.NewUserService().GetBasicUser(p.Context, user.UserID)
		if err != nil {
			return nil, common.FormatError(err)
		}
		user.ToEntity(userInfo)
		return user, nil
	},
}

var addUserInfo = &graphql.Field{
	Description: "新增用户信息",
	Type:        UserInfoType,
	Args:        UserInfoInput,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		user, err := BuildUser(p.Args)
		if err != nil || user.WriteParamCheck() {
			return nil, common.ParamsError
		}
		return nil, common.FormatError(domain.NewUserService().
			ADDBasicUser(p.Context, user.FromEntity()))
	},
}

var updateUserInfo = &graphql.Field{
	Description: "更新用户信息",
	Type:        UserInfoType,
	Args:        UserInfoInput,
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
func BuildUser(args map[string]interface{}) (*User, error) {
	user := User{}
	return &user, mapstructure.Decode(args, &user)
}

// ReadParamCheck true参数校验失败 false参数校验成功
func (u *User) ReadParamCheck() bool {
	if u.UserID == "" || len(u.UserID) > 25 {
		return true
	}
	return false
}

// WriteParamCheck true参数校验失败 false参数校验成功
func (u *User) WriteParamCheck() bool {
	if u.UserID == "" || u.Mobile == "" {
		return true
	}
	return false
}

// FromEntity 从接入层实体转化为业务实体
func (u *User) FromEntity() domain.User {
	return domain.User{
		UserID: u.UserID,
		Mobile: u.Mobile,
		Name:   u.Name,
	}
}

// ToEntity 从业务实体转化为接入层实体
func (u *User) ToEntity(user domain.User) {
	u.UserID = user.UserID
	u.Mobile = user.Mobile
	u.Name = user.Name
}
