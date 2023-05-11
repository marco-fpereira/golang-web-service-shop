package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"web-service/app/util"
)

func DatabaseConnector() *sql.DB {
	connection := "user=postgres dbname=web_shop_db password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	util.ErrorHandler(err)

	return db
}
