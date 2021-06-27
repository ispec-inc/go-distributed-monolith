package main

import (
	"github.com/ispec-inc/monorepo/go/pkg/infra/entity"
	"github.com/ispec-inc/monorepo/go/pkg/rdb"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/config"
	"gorm.io/gorm"
)

var seeds = []interface{}{
	[]*entity.User{
		{
			Name:        "name-1",
			Description: "desc-1",
			Email:       "email-1",
			Articles: []entity.Article{
				{Title: "article-1", Body: "This is article 1."},
				{Title: "article-2", Body: "This is article 2."},
			},
		},
		{
			Name:        "name-2",
			Description: "desc-2",
			Email:       "email-2",
		},
	},
}

func main() {
	var loglev rdb.LogLevel
	if config.MysqlArticle.ShowAllLog {
		loglev = rdb.LogLevelInfo
	} else {
		loglev = rdb.LogLevelError
	}

	db, err := rdb.New(rdb.Config{
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
	if err != nil {
		panic(err)
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, s := range seeds {
			if err := tx.Create(s).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return
	}
}
