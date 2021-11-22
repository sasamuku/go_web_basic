package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwb dbname=gwb password=gwb sslmode=disable")
	if err != nil {
		panic(err)
	}
}
