package main

import (
	"gonews/db"
	"gonews/server"
	// "log"
	// "os"

	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	db.CreateTables()
	server.Run()
}
