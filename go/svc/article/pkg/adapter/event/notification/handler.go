package notification

import (
	"context"
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/yamad07/monorepo/go/pkg/msgbs"
	"github.com/yamad07/monorepo/go/svc/article/pkg/registry"
	"github.com/yamad07/monorepo/go/svc/article/src/v1/notification"
)

type Subscriber struct {
	usecase notification.Usecase
}

func NewSubscriber(rgst registry.Registry) Subscriber {
	usecase := notification.NewUsecase(rgst)
	return Subscriber{usecase}
}

// Print err handling
func (s Subscriber) Do(msg redis.Message) error {
	var m msgbs.Article
	err := json.Unmarshal(msg.Data, &m)
	if err != nil {
		return err
	}

	ipt := notification.Input{
		Title: m.Title,
	}
	err = s.usecase.Send(context.Background(), &ipt)
	if err != nil {
		return err
	}

	return nil
}
