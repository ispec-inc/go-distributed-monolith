package gqlobj

import "github.com/graphql-go/graphql"

var Article = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "",
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "",
			},
			"user": &graphql.Field{
				Type:        User,
				Description: "",
			},
		},
	},
)
