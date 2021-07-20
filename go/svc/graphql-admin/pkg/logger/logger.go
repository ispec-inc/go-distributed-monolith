package logger

import (
	"github.com/ispec-inc/monorepo/go/pkg/applog/logger"
	"github.com/ispec-inc/monorepo/go/pkg/sentry"
	"github.com/ispec-inc/monorepo/go/pkg/stdlog"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/config"
)

var lgr logger.Logger

func Init() (func() error, error) {
	var (
		cleanup func()
		err     error
	)

	switch config.Logger.Type {
	case config.LoggerTypeSentry:
		lgr, cleanup, err = sentry.New(
			sentry.Config{
				Environment: config.Sentry.Env,
				DSN:         config.Sentry.DSN,
				Debug:       config.Sentry.Debug,
			},
		)
		return func() error { cleanup(); return nil }, err
	default:
		lgr = stdlog.New()
		return func() error { return nil }, nil
	}
}

func Get() logger.Logger {
	return lgr
}
