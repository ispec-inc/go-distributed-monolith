package config

import (
	"os"
)

const (
	LoggerTypeStdlog = LoggerType("stdlog")
	LoggerTypeSentry = LoggerType("sentry")
)

type LoggerType string

func init() {
	typ := os.Getenv("ARTICLE_LOGGER_TYPE")
	switch typ {
	case string(LoggerTypeStdlog):
		Logger.Type = LoggerTypeStdlog
	case string(LoggerTypeSentry):
		Logger.Type = LoggerTypeSentry
	default:
		panic("invalid LOGGER_TYPE")
	}
}

var Logger logger

type logger struct {
	Type LoggerType
}
