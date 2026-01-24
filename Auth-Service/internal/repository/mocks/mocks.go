package mocks

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDb() (sqlmock.Sqlmock, *gorm.DB) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	db, mock, _ = sqlmock.New()

	gormDb, _ := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return mock, gormDb
}
