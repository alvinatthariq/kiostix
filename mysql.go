package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// connect - adasdasd
func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/db_book")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
