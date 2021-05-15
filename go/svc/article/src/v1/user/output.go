package user

import (
	"github.com/yamad07/monorepo/go/svc/article/pkg/domain/model"
)

type GetOutput struct {
	User *model.User
}

type CreateOutput struct {
	Users []*model.User
}
