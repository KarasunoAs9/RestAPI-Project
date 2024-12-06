package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connection := "host=localhost user=postgres password=1111 dbname=Rest-App sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connection)

	if err != nil {
		log.Fatal("error with connection to database")
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("no connection with database")
	}

}