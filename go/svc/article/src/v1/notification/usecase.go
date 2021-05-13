package notification

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/yamad07/monorepo/go/pkg/applog"
	"github.com/yamad07/monorepo/go/svc/article/pkg/registry"
)

type Usecase struct {
	log applog.AppLog
}

func NewUsecase(rgst registry.Registry) Usecase {
	return Usecase{
		log: applog.New(rgst.Logger().New()),
	}
}

func (u Usecase) Send(ctx context.Context, ipt *Input) error {
	spew.Dump(ipt)
	return nil
}
