package rdb

import (
	"time"

	"gorm.io/gorm/logger"
)

const (
	LogLevelInfo  = LogLevel(logger.Info)
	LogLevelError = LogLevel(logger.Error)

	DBMSMySQL      = DBMS("mysql")
	DBMSPostgreSQL = DBMS("postgres")
)

type (
	Connection struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}

	Config struct {
		DBMS        DBMS
		Conn        Connection
		LogLevel    LogLevel
		MaxIdleConn int
		MaxOpenConn int
		MaxLifetime time.Duration
	}

	TestConfig struct {
		DBMS DBMS
		Conn Connection
	}

	LogLevel logger.LogLevel

	DBMS string
)
