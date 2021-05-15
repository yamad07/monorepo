package event

import (
	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/svc/article/pkg/adapter/event/notification"
	"github.com/yamad07/monorepo/go/svc/article/pkg/registry"
)

func NewSubscriber(bs msgbs.MessageBus) (msgbs.Subscriber, func() error, error) {
	repo, repoCleanup, err := registry.NewRepository()
	if err != nil {
		return msgbs.Subscriber{}, nil, err
	}

	lgr, lgrCleanup, err := registry.NewLogger()
	if err != nil {
		return msgbs.Subscriber{}, nil, err
	}
	cleanup := func() error {
		repoCleanup()
		lgrCleanup()
		return nil
	}

	rgst := registry.NewRegistry(repo, lgr)
	r := msgbs.NewSubscriber(bs)

	subsc := notification.NewSubscriber(rgst)

	r.Subscribe(msgbs.AddArticle, subsc.Notify)

	return r, cleanup, nil
}
