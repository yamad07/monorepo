package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yamad07/monorepo/go/pkg/config"
	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	HTTPHandler     http.Handler
	SubscribeRouter *msgbs.Router
}

func NewServer(
	h http.Handler,
	r *msgbs.Router,
) Server {
	return Server{
		HTTPHandler:     h,
		SubscribeRouter: r,
	}
}

func (s Server) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		port := fmt.Sprintf(":%d", config.Router.Port)
		err := http.ListenAndServe(port, s.HTTPHandler)
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	g.Go(func() error {
		s.SubscribeRouter.Serve()
		return nil
	})

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case <-ctx.Done():
		break
	case <-cs:
		break
	}

	_, tcancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer tcancel()

	err := g.Wait()
	if err != nil {
		os.Exit(2)
	}
}
