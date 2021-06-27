package rdb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MysqlDSN    = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true"
	PostgresDSN = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	ErrStrUnkownDBMS = "unknown DBMS: %s"
)

func dialector(dbms DBMS, c Connection) (gorm.Dialector, error) {
	dsn, err := dsn(dbms, c)
	if err != nil {
		return nil, err
	}
	switch dbms {
	case DBMSMySQL:
		return mysql.Open(dsn), nil
	case DBMSPostgreSQL:
		return postgres.Open(dsn), nil
	default:
		return nil, fmt.Errorf(ErrStrUnkownDBMS, string(dbms))
	}
}

func dsn(dbms DBMS, c Connection) (string, error) {
	switch dbms {
	case DBMSMySQL:
		return fmt.Sprintf(MysqlDSN, c.User, c.Password, c.Host, c.Port, c.Database), nil
	case DBMSPostgreSQL:
		return fmt.Sprintf(PostgresDSN, c.User, c.Password, c.Host, c.Port, c.Database), nil
	default:
		return "", fmt.Errorf(ErrStrUnkownDBMS, string(dbms))
	}
}
