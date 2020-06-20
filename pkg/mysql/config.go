package mysql

import (
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func SetDB() error {
	DBMS := os.Getenv("DB_TYPE")
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	NAME := os.Getenv("DB_NAME")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")

	CONNECT := USER + ":" + PASS + "@(" + HOST + ":" + PORT + ")/" + NAME + "?charset=utf8mb4&parseTime=true"

	db_, err := gorm.Open(DBMS, CONNECT)

	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPENC_CONNS"))
	maxLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))

	db_.DB().SetMaxIdleConns(maxIdle)
	db_.DB().SetMaxOpenConns(maxOpen)
	db_.DB().SetConnMaxLifetime(time.Duration(maxLifetime))

	db = db_
	return err
}

func GetConnection() *gorm.DB {
	return db
}
