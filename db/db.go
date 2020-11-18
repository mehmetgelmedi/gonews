package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateTables() {
	db := Connect()
	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS public.user (id SERIAL PRIMARY KEY, username VARCHAR(20) UNIQUE, password VARCHAR(60), is_admin boolean)")
	if err != nil {
		panic(err)
	}
}

func Connect() *sql.DB {
	conninfo := fmt.Sprintf("postgresql://%s:%s@postgres/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Panic(err)
	}
	return db
}
