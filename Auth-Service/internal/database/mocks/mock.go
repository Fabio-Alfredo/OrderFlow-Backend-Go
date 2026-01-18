package mocks

import (
	"Auth-Service/pkg/config"
	"database/sql"
	"fmt"
)

func GetDbConnectionMock(config config.IConfig) *sql.DB {

	dns := GetDnsMock(config)
	db, _ := sql.Open(config.GetString("datasource.driver"), dns)
	return db
}

func GetDnsMock(config config.IConfig) string {
	username := config.GetString("datasource.username")
	password := config.GetString("datasource.password")
	host := config.GetString("datasource.host")
	database := config.GetString("datasource.database")
	zone := config.GetString("datasource.zone")
	port := config.GetString("datasource.port")

	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True&loc=%s", username, password, port, host, database, zone)
}
