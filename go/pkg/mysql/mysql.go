package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	LogLevelInfo  = LogLevel(logger.Info)
	LogLevelError = LogLevel(logger.Error)
)

type (
	Config struct {
		User        string
		Password    string
		Host        string
		Port        string
		Database    string
		LogLevel    LogLevel
		MaxIdleConn int
		MaxOpenConn int
		MaxLifetime time.Duration
	}
	LogLevel logger.LogLevel
)

func New(c Config) (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.Database,
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(c.LogLevel)),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(c.MaxLifetime)

	return db, nil
}
