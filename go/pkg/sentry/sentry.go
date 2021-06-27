package sentry

import (
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/ispec-inc/monorepo/go/pkg/applog/logger"
)

type Logger struct{}

type Config struct {
	Environment string
	DSN         string
	Debug       bool
}

func New(c Config) (*Logger, func(), error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         c.DSN,
		Environment: c.Environment,
		Debug:       c.Debug,
	})
	cleanup := func() { sentry.Flush(2 * time.Second) }
	return &Logger{}, cleanup, err
}

func (l *Logger) Error(user logger.User, err logger.Error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		event := sentry.NewEvent()

		event.Level = sentry.LevelError
		event.User = sentry.User{
			ID:       user.ID,
			Username: user.Name,
		}
		event.Tags = map[string]string{"code": err.Code}
		event.Extra = map[string]interface{}{"message": err.Message}

		event.Exception = []sentry.Exception{{
			Value:      err.Message,
			Type:       err.ErrorType,
			Stacktrace: sentry.ExtractStacktrace(err.Error),
		}}

		sentry.CaptureEvent(event)
	})
}
