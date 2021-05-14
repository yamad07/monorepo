package article

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/monorepo/go/pkg/msgbs"
)

func New(bs msgbs.MessageBus) http.Handler {
	r := chi.NewRouter()
	h := newController(bs)

	r.Get("/", h.list)
	r.Get("/{id}", h.get)
	r.Post("/", h.create)
	r.Put("/{id}", h.update)
	r.Delete("/{id}", h.delete)
	return r
}
