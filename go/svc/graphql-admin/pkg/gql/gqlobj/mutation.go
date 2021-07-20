package gqlobj

import "github.com/graphql-go/graphql"

type MutationResolverRegistry interface {
	CreateArticle() func(params graphql.ResolveParams) (interface{}, error)
}

func Mutation(r MutationResolverRegistry) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createArticle": &graphql.Field{
				Type:        Article,
				Description: "",
				Resolve:     r.CreateArticle(),
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "",
					},
					"description": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "",
					},
				},
			},
		},
	})
}
