package rdb

import (
	"fmt"

	txdb "github.com/DATA-DOG/go-txdb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const TestDriver = "txdb"

func InitTest(c TestConfig) error {
	dsn, err := dsn(c.DBMS, c.Conn)
	if err != nil {
		return err
	}
	txdb.Register(TestDriver, "mysql", dsn)
	return nil
}

func NewTest(dbms DBMS, name string) (*gorm.DB, error) {
	dialector, err := testDialector(dbms, name)
	if err != nil {
		return nil, err
	}
	return gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func testDialector(dbms DBMS, name string) (gorm.Dialector, error) {
	switch dbms {
	case DBMSMySQL:
		return mysql.New(mysql.Config{
			DriverName: TestDriver,
			DSN:        name,
		}), nil
	case DBMSPostgreSQL:
		return postgres.New(postgres.Config{
			DriverName: TestDriver,
			DSN:        name,
		}), nil
	default:
		return nil, fmt.Errorf(ErrStrUnkownDBMS, string(dbms))
	}
}
