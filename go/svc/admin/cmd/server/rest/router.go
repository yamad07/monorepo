package rest

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/pkg/presenter"
	v1 "github.com/yamad07/monorepo/go/svc/admin/pkg/controller/rest/v1"
	"github.com/yamad07/monorepo/go/svc/admin/pkg/database"
	"github.com/yamad07/monorepo/go/svc/admin/pkg/logger"
)

func NewRouter(pubsub msgbs.RedisPubSub) (http.Handler, func() error, error) {
	if err := database.Init(nil); err != nil {
		return nil, nil, err
	}
	cleanup, err := logger.Init()
	if err != nil {
		return nil, nil, err
	}

	r := chi.NewRouter()
	r = commonMiddleware(r)

	r.Mount("/v1", v1.NewRouter(pubsub))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Response(w, map[string]string{"messsage": "ok"})
	})

	return r, cleanup, nil
}
