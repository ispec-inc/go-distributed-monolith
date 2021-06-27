package image

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/registry"
)

func NewRouter(rgst registry.Registry) http.Handler {
	h := newHandler(rgst)
	r := chi.NewRouter()
	r.Post("/", h.Create)
	return r
}
