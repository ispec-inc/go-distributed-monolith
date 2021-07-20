package database

import (
	"gorm.io/gorm"

	"github.com/ispec-inc/monorepo/go/pkg/rdb"
	"github.com/ispec-inc/monorepo/go/svc/graphql-admin/pkg/config"
)

var db *gorm.DB

func Init(indb *gorm.DB) (err error) {
	if indb != nil {
		db = indb
		return nil
	}

	var loglev rdb.LogLevel
	if config.MysqlArticle.ShowAllLog {
		loglev = rdb.LogLevelInfo
	} else {
		loglev = rdb.LogLevelError
	}

	db, err = rdb.New(rdb.Config{
		DBMS: rdb.DBMSMySQL,
		Conn: rdb.Connection{
			User:     config.MysqlArticle.User,
			Password: config.MysqlArticle.Password,
			Host:     config.MysqlArticle.Host,
			Port:     config.MysqlArticle.Port,
			Database: config.MysqlArticle.Database,
		},
		LogLevel:    loglev,
		MaxIdleConn: config.MysqlArticle.MaxIdleConn,
		MaxOpenConn: config.MysqlArticle.MaxOpenConn,
		MaxLifetime: config.MysqlArticle.MaxLifetime,
	})
	return err
}

func Get() *gorm.DB {
	return db
}
