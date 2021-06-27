package config

import "github.com/caarlos0/env/v6"

func init() {
	if err := env.Parse(&Sentry); err != nil {
		panic(err)
	}
}

var Sentry sentry

type sentry struct {
	DSN   string `env:"NOTIFICATION_SENTRY_DSN"`
	Env   string `env:"NOTIFICATION_SENTRY_ENVIRONMENT"`
	Debug bool   `env:"NOTIFICATION_SENTRY_DEBUG"`
}
