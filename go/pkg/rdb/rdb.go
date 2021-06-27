package rdb

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(c Config) (*gorm.DB, error) {
	dlctr, err := dialector(c.DBMS, c.Conn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dlctr, &gorm.Config{
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
