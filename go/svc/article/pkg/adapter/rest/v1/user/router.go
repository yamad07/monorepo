package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/monorepo/go/svc/article/pkg/registry"
)

func NewRouter(rgst registry.Registry) http.Handler {
	r := chi.NewRouter()
	h := NewHandler(rgst)

	r.Get("/{id}", h.Get)
	r.Post("/", h.Create)
	return r
}
