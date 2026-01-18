package database

import (
	"Auth-Service/pkg/config"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dnsFormat = "%s:%s@(%s:%s)/%s?parseTime=True&loc=%s"

type sqlConfig struct {
	config config.IConfig
	db     *gorm.DB
}

func NewSQLConfig(config config.IConfig) ISqlDB {
	return &sqlConfig{
		config: config,
	}
}

func (s *sqlConfig) GetDB() (*gorm.DB, error) {

	if s.db != nil {
		return s.db, nil
	}

	sqlDb, sqlErr := s.openConnection()
	if sqlErr != nil {
		return nil, sqlErr
	}

	gormDb, gormErr := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if gormErr != nil {
		return nil, gormErr
	}

	s.db = gormDb
	return gormDb, nil
}

func (s *sqlConfig) openConnection() (*sql.DB, error) {
	driver := s.config.GetString("datasource.driver")
	dns := s.getDns()

	sqlDb, sqlErr := sql.Open(driver, dns)
	if sqlErr != nil {
		return nil, sqlErr
	}

	return sqlDb, nil
}

func (s *sqlConfig) getDns() string {
	username := s.config.GetString("datasource.username")
	password := s.config.GetString("datasource.password")
	host := s.config.GetString("datasource.host")
	database := s.config.GetString("datasource.database")
	zone := s.config.GetString("datasource.zone")
	port := s.config.GetString("datasource.port")

	return fmt.Sprintf(dnsFormat, username, password, port, host, database, zone)
}
