package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&MysqlArticle); err != nil {
		panic(err)
	}
}

var MysqlArticle mysqlArticle

type mysqlArticle struct {
	User        string        `env:"ARTICLE_MYSQL_ARTICLE_USER"`
	Password    string        `env:"ARTICLE_MYSQL_ARTICLE_PASSWORD"`
	Database    string        `env:"ARTICLE_MYSQL_ARTICLE_DATABASE"`
	Host        string        `env:"ARTICLE_MYSQL_ARTICLE_HOST"`
	Port        string        `env:"ARTICLE_MYSQL_ARTICLE_PORT"`
	ShowAllLog  bool          `env:"ARTICLE_MYSQL_ARTICLE_SHOW_ALL_LOG"`
	MaxIdleConn int           `env:"ARTICLE_MYSQL_ARTICLE_MAX_IDLE_CONN"`
	MaxOpenConn int           `env:"ARTICLE_MYSQL_ARTICLE_MAX_OPEN_CONN"`
	MaxLifetime time.Duration `env:"ARTICLE_MYSQL_ARTICLE_MAX_CONN_LIFETIME"`
}
