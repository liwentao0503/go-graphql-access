package api

import (
	"github.com/graphql-go/graphql"
)

// QueryType api查询声明
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"getUserInfo": getUserInfo,
	},
})

// MutationType 增、改声明
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"addUser":    addUserInfo,
		"updateUser": updateUserInfo,
	},
})

// SubscriptionType 订阅声明
var SubscriptionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "subscription",
	Fields: graphql.Fields{
		"flowerUser": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
	},
})
