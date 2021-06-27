package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&Router); err != nil {
		panic(err)
	}
}

var Router router

type router struct {
	Timeout      time.Duration `env:"ADMIN_ROUTER_TIMEOUT"`
	AllowOrigins []string      `env:"ADMIN_ROUTER_ALLOW_ORIGINS" envSeparator:","`
}
