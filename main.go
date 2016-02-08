package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
)

var vehicleStateEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "VehicleStatus",
	Values: graphql.EnumValueConfigMap{
		"FREE":     &graphql.EnumValueConfig{Value: "FREE", Description: "The vehicle is free for usage"},
		"RESERVED": &graphql.EnumValueConfig{Value: "RESERVED", Description: "The vehicle is reserved"},
	},
})

var vehicleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Vehicle",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String, Description: "Vehicle id"},
		"name":  &graphql.Field{Type: graphql.String, Description: "Vehicle name"},
		"state": &graphql.Field{Type: vehicleStateEnum, Description: "true=FREE, false=RESERVED"},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Vehicle": &graphql.Field{
			Description: "Get a vehicle",
			Type:        vehicleType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String, Description: "Vehicle identifier"},
			},
			Resolve: findVehicleById,
		},
	},
})

func findVehicleById(params graphql.ResolveParams) (interface{}, error) {
	// About params...
	// https://github.com/graphql-go/graphql/blob/26c58bd73a3cf9177fb534f57c9698a7bc6277a2/definition.go
	//   type ResolveParams struct {
	//     Source interface{}
	//     Args   map[string]interface{}
	//     Info   ResolveInfo
	//     Schema Schema
	//     //This can be used to provide per-request state
	//     //from the application.
	//     Context context.Context
	//   }
	id, _ := params.Args["id"]

	fmt.Printf("let's find a vehicle with id=%s...\n", id.(string))
	//return map[string] interface{} {"name": "Toyota " + params.Args["id"], "state": vehicleStateEnum.Values()["FREE"]},
	vehicleName := "Toyota " + id.(string)
	return map[string]interface{}{"name": vehicleName, "state": "FREE"}, nil
}

func main() {

	var schema, err = graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
	if err != nil {
		panic(err)
	}

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
