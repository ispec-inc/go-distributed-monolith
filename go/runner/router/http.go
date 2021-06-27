package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/monorepo/go/pkg/presenter"
	admin "github.com/ispec-inc/monorepo/go/svc/admin/runner/router"
	article "github.com/ispec-inc/monorepo/go/svc/article/runner/router"
	media "github.com/ispec-inc/monorepo/go/svc/media/runner/router"
	"go.uber.org/multierr"
)

const PORT = 9000

func NewHTTP() (*http.Server, func() error, error) {
	routers := []struct {
		new  func() (http.Handler, func() error, error)
		path string
	}{
		{new: admin.NewREST, path: "/admin"},
		{new: article.NewREST, path: "/article"},
		{new: media.NewREST, path: "/media"},
	}

	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Response(w, map[string]string{"messsage": "ok"})
	})

	clnups := make([]func() error, len(routers))
	for i, router := range routers {
		h, clnup, err := router.new()
		if err != nil {
			return nil, nil, err
		}
		r.Mount(router.path, h)
		clnups[i] = clnup
	}

	clnup := func() error {
		var errs error
		for _, clnup := range clnups {
			errs = multierr.Append(errs, clnup())
		}
		return errs
	}

	port := fmt.Sprintf(":%d", PORT)
	srv := &http.Server{Addr: port, Handler: r}
	return srv, clnup, nil

}
