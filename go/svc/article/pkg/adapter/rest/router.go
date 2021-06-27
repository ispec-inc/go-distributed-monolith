package rest

import (
	"net/http"

	"github.com/go-chi/chi"

	v1 "github.com/ispec-inc/monorepo/go/svc/article/pkg/adapter/rest/v1"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/registry"
)

func NewRouter(rgst registry.Registry) http.Handler {
	r := chi.NewRouter()

	r.Mount("/v1", v1.NewRouter(rgst))

	return r
}
