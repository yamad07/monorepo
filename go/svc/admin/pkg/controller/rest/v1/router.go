package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/svc/admin/pkg/controller/rest/v1/article"
)

func NewRouter(bs msgbs.MessageBus) http.Handler {
	r := chi.NewRouter()

	r.Mount("/articles", article.New(bs))

	return r
}
