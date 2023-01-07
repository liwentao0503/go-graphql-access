package api

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"getUserInfo": getUserInfo,
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"addUser":    addUserInfo,
		"updateUser": updateUserInfo,
	},
})

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
