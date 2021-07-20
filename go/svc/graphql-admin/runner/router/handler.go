package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/ispec-inc/monorepo/go/pkg/presenter"
)

type graphqlHandler struct {
	schema graphql.Schema
}

type request struct {
	Query string `json:"query"`
}

func (h graphqlHandler) run(w http.ResponseWriter, r *http.Request) {
	q := &request{}
	if err := json.NewDecoder(r.Body).Decode(q); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	result := h.executeQuery(q.Query, h.schema)
	presenter.Response(w, result)
}

func (h graphqlHandler) executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}

	return result
}
