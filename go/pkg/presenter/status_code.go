package presenter

import (
	"net/http"

	"github.com/yamad07/monorepo/go/pkg/apperror"
)

var code2status = map[apperror.Code]int{
	apperror.CodeError:        http.StatusInternalServerError,
	apperror.CodeInvalid:      http.StatusBadRequest,
	apperror.CodeUnauthorized: http.StatusUnauthorized,
	apperror.CodeNotFound:     http.StatusNotFound,
}
