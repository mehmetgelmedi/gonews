package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connect() *sql.DB {
	conninfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Panic(err)
	}
	return db
}