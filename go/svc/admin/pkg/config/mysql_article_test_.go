package config

import (
	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&MysqlArticleTest); err != nil {
		panic(err)
	}
}

var MysqlArticleTest mysqlArticleTest

type mysqlArticleTest struct {
	User     string `env:"ADMIN_MYSQL_ARTICLE_TEST_USER"`
	Password string `env:"ADMIN_MYSQL_ARTICLE_TEST_PASSWORD"`
	Database string `env:"ADMIN_MYSQL_ARTICLE_TEST_DATABASE"`
	Host     string `env:"ADMIN_MYSQL_ARTICLE_TEST_HOST"`
	Port     string `env:"ADMIN_MYSQL_ARTICLE_TEST_PORT"`
}
