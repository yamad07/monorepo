package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/ispec-inc/monorepo/go/pkg/config"
)

func commonMiddleware(r *chi.Mux) *chi.Mux {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(config.Router.Timeout))
	return r
}
