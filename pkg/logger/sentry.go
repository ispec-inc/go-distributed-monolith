package logger

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
)

type sentryLogger struct{}

func newSentryLogger(options SentryOptions) (*sentryLogger, func(), error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         options.DSN,
		Environment: options.Environment,
		Debug:       options.Debug,
	})
	cleanup := func() { sentry.Flush(2 * time.Second) }
	return &sentryLogger{}, cleanup, err
}

func (r *sentryLogger) SetUser(ctx context.Context, id, name string) context.Context {
	ctx = context.WithValue(ctx, userIDKey, id)
	ctx = context.WithValue(ctx, userNameKey, name)
	return ctx
}

func (r *sentryLogger) Error(ctx context.Context, code, message string, err error) {
	var uid, uname string
	if v := ctx.Value(userIDKey); v != nil {
		uid = v.(string)
	}
	if v := ctx.Value(userNameKey); v != nil {
		uname = v.(string)
	}

	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{
			ID:       uid,
			Username: uname,
		})
		scope.SetTags(map[string]string{
			"code": code,
		})
		scope.SetExtra("message", message)
		sentry.CaptureException(err)
	})
}
