package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"

	"github.com/yamad07/monorepo/go/pkg/config"
	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/pkg/presenter"
	admin_rest "github.com/yamad07/monorepo/go/svc/admin/cmd/server/rest"
	article_event "github.com/yamad07/monorepo/go/svc/article/cmd/server/event"
	article_rest "github.com/yamad07/monorepo/go/svc/article/cmd/server/rest"
)

func main() {
	msgbs.Init(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	ps := msgbs.NewRedis()
	// TODO pointer
	adh, adclnup, err := admin_rest.NewRouter(*ps)
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

	atclevnt, evntclnup, err := article_event.NewRouter()
	if err != nil {
		panic(err)
	}
	defer evntclnup()

	sr := msgbs.NewRouter()
	sr.Mount(atclevnt)

	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := http.ListenAndServe(port, r); err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	g.Go(func() error {
		// TODO error handling
		sr.Serve()
		return nil
	})

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	err = g.Wait()
	if err != nil {
		log.Printf("server returning an error %v\n", err)
		os.Exit(2)
	}

	log.Println("all servers are stopped")
}
