package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbname     = os.Getenv("LIFETRAPS_DBNAME")
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@/"+dbname)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database.")
	return db
}
