package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionUrl := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", connectionUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection open")

	lines, err := db.Query("SELECT * FROM users")

	if err == nil {
		log.Fatal(err)
	}

	defer lines.Close()

	fmt.Println(lines)
}
