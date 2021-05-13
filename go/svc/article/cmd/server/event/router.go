package event

import (
	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/svc/article/pkg/adapter/event/notification"
	"github.com/yamad07/monorepo/go/svc/article/pkg/registry"
)

func NewRouter() (msgbs.Router, func() error, error) {
	repo, repoCleanup, err := registry.NewRepository()
	if err != nil {
		return msgbs.Router{}, nil, err
	}

	lgr, lgrCleanup, err := registry.NewLogger()
	if err != nil {
		return msgbs.Router{}, nil, err
	}
	cleanup := func() error {
		repoCleanup()
		lgrCleanup()
		return nil
	}

	rgst := registry.NewRegistry(repo, lgr)
	r := msgbs.NewRouter()

	subsc := notification.NewSubscriber(rgst)
	r.Subscribe(msgbs.AddArticle, subsc)

	return r, cleanup, nil
}
