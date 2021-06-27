package middleware

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type CommonConfig struct {
	Timeout      time.Duration
	AllowOrigins []string
}

func Common(r *chi.Mux, c CommonConfig) *chi.Mux {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(c.Timeout))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: c.AllowOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept", "Accept-Language", "Authorization",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return r
}
