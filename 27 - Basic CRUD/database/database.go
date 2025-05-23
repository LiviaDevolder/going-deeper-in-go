package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Connection driver with MySQL
)

// Open connection with db
func Connect() (*sql.DB, error) {
	connectionUrl := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionUrl)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
