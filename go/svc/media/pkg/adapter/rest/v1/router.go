package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/adapter/rest/v1/image"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/registry"
)

func NewRouter(rgst registry.Registry) http.Handler {
	r := chi.NewRouter()

	r.Mount("/images", image.NewRouter(rgst))

	return r
}
