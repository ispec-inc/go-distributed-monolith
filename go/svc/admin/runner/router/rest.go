package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ispec-inc/monorepo/go/pkg/middleware"
	"github.com/ispec-inc/monorepo/go/pkg/presenter"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/config"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/controller/rest"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/database"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/logger"
	"github.com/ispec-inc/monorepo/go/svc/admin/pkg/redisbs"
)

func NewREST() (http.Handler, func() error, error) {
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

	r := chi.NewRouter()
	r = middleware.Common(r, middleware.CommonConfig{
		Timeout:      config.Router.Timeout,
		AllowOrigins: config.Router.AllowOrigins,
	})

	r.Mount("/", rest.NewRouter())

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Response(w, map[string]string{"messsage": "ok"})
	})

	return r, cleanup, nil
}
