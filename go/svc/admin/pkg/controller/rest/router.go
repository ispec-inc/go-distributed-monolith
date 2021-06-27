package rest

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "github.com/ispec-inc/monorepo/go/svc/admin/pkg/controller/rest/v1"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Mount("/v1", v1.NewRouter())

	return r
}
