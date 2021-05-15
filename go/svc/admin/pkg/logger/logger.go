package logger

import (
	"github.com/yamad07/monorepo/go/pkg/applog/logger"
	"github.com/yamad07/monorepo/go/pkg/config"
	"github.com/yamad07/monorepo/go/pkg/sentry"
	"github.com/yamad07/monorepo/go/pkg/stdlog"
)

var lgr logger.Logger

func Init() (func() error, error) {
	var (
		cleanup func()
		err     error
	)

	switch config.App.Env {
	case config.EnvDev:
		lgr = stdlog.New()
		return func() error { return nil }, nil
	default:
		lgr, cleanup, err = sentry.New(
			sentry.Options{
				Environment: config.Sentry.DSN,
				DSN:         config.Sentry.DSN,
				Debug:       config.Sentry.Debug,
			},
		)
		return func() error { cleanup(); return nil }, err
	}
}

func Get() logger.Logger {
	return lgr
}
