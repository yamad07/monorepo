package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/monorepo/go/svc/article/pkg/adapter/rest/v1/user"
	"github.com/yamad07/monorepo/go/svc/article/pkg/registry"
)

func NewRouter(rgst registry.Registry) http.Handler {
	r := chi.NewRouter()

	r.Mount("/users", user.NewRouter(rgst))

	return r
}
