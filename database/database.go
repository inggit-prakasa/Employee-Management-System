package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/inggit_prakasa/Employee/config"
)

var database *sql.DB
var err error

func init() {
	conf := config.GetConfig()

	conString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	database, err = sql.Open("mysql", conString)

	if err != nil {
		panic("connection Error")
	}

	err = database.Ping()
	if err != nil {
		panic("DNS salah")
	}

}

func Connection() *sql.DB {
	return database
}
