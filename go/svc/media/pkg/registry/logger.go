package registry

import (
	"github.com/ispec-inc/monorepo/go/pkg/applog/logger"
	"github.com/ispec-inc/monorepo/go/pkg/sentry"
	"github.com/ispec-inc/monorepo/go/pkg/stdlog"
	"github.com/ispec-inc/monorepo/go/svc/media/pkg/config"
)

type Logger struct {
	lgr logger.Logger
}

func NewLogger() (Logger, func() error, error) {
	var (
		lgr     logger.Logger
		clenaup func() error
		err     error
	)

	switch config.Logger.Type {
	case config.LoggerTypeSentry:
		slgr, scleanup, serr := sentry.New(
			sentry.Config{
				Environment: config.Sentry.Env,
				DSN:         config.Sentry.DSN,
				Debug:       config.Sentry.Debug,
			},
		)
		lgr = slgr
		clenaup = func() error { scleanup(); return nil }
		err = serr
	default:
		lgr = stdlog.New()
	}

	return Logger{lgr}, clenaup, err
}

func (l Logger) New() logger.Logger {
	return l.lgr
}
