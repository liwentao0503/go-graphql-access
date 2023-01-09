package main

import (
	"net/http"

	"code_struct/api"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:        api.QueryType,
	Mutation:     api.MutationType,
	Subscription: api.SubscriptionType,
})

// create a graphl-go HTTP handler with our previously defined schema
// and we also set it to return pretty JSON output
var handler *gqlhandler.Handler = gqlhandler.New(&gqlhandler.Config{
	Schema:   &Schema,
	Pretty:   true,
	GraphiQL: true,
})

var router = map[string]*gqlhandler.Handler{
	"/graphql":   handler,
	"/graphqlv2": handler,
}

func main() {
	Init()

	for k, v := range router {
		http.Handle(k, v)
	}

	// and serve!
	http.ListenAndServe(":8080", nil)
}

func Init() {
	// mysql.InitDB()
}
