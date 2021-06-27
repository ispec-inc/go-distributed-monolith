package registry

import (
	"gorm.io/gorm"

	"github.com/ispec-inc/monorepo/go/pkg/rdb"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/config"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/domain/repository"
	"github.com/ispec-inc/monorepo/go/svc/article/pkg/infra/dao"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (Repository, func() error, error) {
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
		return Repository{}, func() error { return nil }, err
	}

	repo := Repository{
		db: db,
	}
	cleanup := func() error { return nil }
	return repo, cleanup, nil
}

func (repo Repository) NewUser() repository.User {
	return dao.NewUser(repo.db)
}
