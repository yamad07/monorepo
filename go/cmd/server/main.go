package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/yamad07/monorepo/go/pkg/config"
	"github.com/yamad07/monorepo/go/pkg/presenter"
	admin_rest "github.com/yamad07/monorepo/go/svc/admin/cmd/server/rest"
	article_rest "github.com/yamad07/monorepo/go/svc/article/cmd/server/rest"
)

func main() {
	adh, adclnup, err := admin_rest.NewRouter()
	if err != nil {
		panic(err)
	}
	defer adclnup()

	atclh, arclclnup, err := article_rest.NewRouter()
	if err != nil {
		panic(err)
	}
	defer arclclnup()

	r := chi.NewRouter()

	r.Mount("/admin", adh)
	r.Mount("/articles", atclh)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Response(w, map[string]string{"messsage": "ok"})
	})

	port := fmt.Sprintf(":%d", config.Router.Port)
	http.ListenAndServe(port, r)
}
