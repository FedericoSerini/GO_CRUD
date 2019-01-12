package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDatabaseConnection() (*sql.DB, error) {
	DatabaseConnection, err := sql.Open("mysql",
		"root:root@tcp(localhost:3306)/test?charset=utf8")

	if err != nil || DatabaseConnection.Ping() != nil {
		log.Println("ERROR: DATABASE_INIT: ", err)
	}

	return DatabaseConnection, err
}
