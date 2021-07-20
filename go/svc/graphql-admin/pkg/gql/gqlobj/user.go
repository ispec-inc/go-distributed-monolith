package gqlobj

import "github.com/graphql-go/graphql"

var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "",
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "",
			},
		},
	},
)
