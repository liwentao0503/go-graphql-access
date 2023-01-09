package router

import (
	"go-graphql-access/api"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var Router = map[string]*gqlhandler.Handler{
	"/graphql":   handler,
	"/graphqlv2": handler,
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:        api.QueryType,
	Mutation:     api.MutationType,
	Subscription: api.SubscriptionType,
})

// create a graphl-go HTTP handler with our previously defined schema
// and we also set it to return pretty JSON output
var handler *gqlhandler.Handler = gqlhandler.New(&gqlhandler.Config{
	Schema:   &schema,
	Pretty:   true,
	GraphiQL: true,
})
