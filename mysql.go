package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/gotest")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
