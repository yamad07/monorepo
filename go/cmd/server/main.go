package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/pkg/presenter"
	"github.com/yamad07/monorepo/go/pkg/redis"
	admin_rest "github.com/yamad07/monorepo/go/svc/admin/cmd/server/rest"
	article_event "github.com/yamad07/monorepo/go/svc/article/cmd/server/event"
	article_rest "github.com/yamad07/monorepo/go/svc/article/cmd/server/rest"
)

func main() {

	bs, err := msgbsConn()
	if err != nil {
		panic(err)
	}

	hr, hclnup, err := httpRouter(bs)
	if err != nil {
		panic(err)
	}
	defer hclnup()

	sr, sclnup, err := subscribeRouter(bs)
	if err != nil {
		panic(err)
	}
	defer sclnup()

	ctx := context.Background()

	server := NewServer(hr, sr)
	server.Run(ctx)
}

func msgbsConn() (msgbs.MessageBus, error) {
	rcon, err := redis.New()
	if err != nil {
		return nil, err
	}

	pscon, err := redis.NewPubSub()
	if err != nil {
		return nil, err
	}

	bs := msgbs.NewRedis(pscon, &rcon)

	return bs, nil
}

func subscribeRouter(bs msgbs.MessageBus) (*msgbs.Router, func() error, error) {
	atclevnt, evntclnup, err := article_event.NewRouter()
	if err != nil {
		return nil, nil, err
	}

	sr := msgbs.NewRouter(bs)
	sr.Mount(atclevnt)
	return &sr, evntclnup, nil
}

func httpRouter(bs msgbs.MessageBus) (http.Handler, func() error, error) {
	adh, adclnup, err := admin_rest.NewRouter(bs)
	if err != nil {
		return nil, nil, err
	}

	atclh, arclclnup, err := article_rest.NewRouter()
	if err != nil {
		return nil, nil, err
	}

	r := chi.NewRouter()

	r.Mount("/admin", adh)
	r.Mount("/articles", atclh)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Response(w, map[string]string{"messsage": "ok"})
	})

	clnup := func() error {
		err := adclnup()
		if err != nil {
			return err
		}

		err = arclclnup()
		if err != nil {
			return err
		}
		return nil
	}

	return r, clnup, nil

}
