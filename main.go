package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Description: "Latest post field",
			Type: graphql.String,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// About params... 
				// https://github.com/graphql-go/graphql/blob/26c58bd73a3cf9177fb534f57c9698a7bc6277a2/definition.go
				return "Hello World!", nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	// create a graphql-go HTTP hanlder with our prevously defined schema
	// and we also set it to return pretty JSON outpput
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at '/graphql' and serve!
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
