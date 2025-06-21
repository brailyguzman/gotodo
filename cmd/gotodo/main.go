package main

import (
	"log"
	"os"

	"github.com/brailyguzman/gotodo/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("POSTGRES_DSN")

	if dsn == "" {
		log.Fatal("POSTGRES_DSN environment variable is not set")
	}

	db.ConnectDatabase(dsn)
}
