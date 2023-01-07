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

func main() {
	Init()

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)
	http.Handle("/graphqlv2", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}

func Init() {
	// mysql.InitDB()
}
