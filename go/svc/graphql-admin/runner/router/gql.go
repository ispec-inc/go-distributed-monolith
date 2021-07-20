package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/database"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/gql/gqlobj"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/gql/resolver"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/logger"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/redisbs"
)

func NewGraphQL() (http.Handler, func() error, error) {
	if err := database.Init(nil); err != nil {
		return nil, nil, err
	}
	cleanup, err := logger.Init()
	if err != nil {
		return nil, nil, err
	}

	if err := redisbs.Init(); err != nil {
		return nil, nil, err
	}

	r, err := initHandler()
	if err != nil {
		return nil, nil, err
	}

	return r, cleanup, nil
}

func initHandler() (http.Handler, error) {
	schemaConfig := graphql.SchemaConfig{
		Query: gqlobj.Query(resolver.NewQueryResolver()),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	h := graphqlHandler{
		schema: schema,
	}
	r := chi.NewRouter()
	r.Post("/graphql", h.run)

	return r, nil
}
