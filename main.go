package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=postgres sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}


	rows, err := db.Query("SELECT version()")
	if err != nil {
		panic(err)
	}

	println(rows)
}