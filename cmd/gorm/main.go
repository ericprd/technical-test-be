package main

import (
	"log"

	"github.com/ericprd/technical-test/config"
	"github.com/ericprd/technical-test/database"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file, continuing without it")
	}

	config.InitConfig()
}

func main() {
	db, err := database.New()

	if err != nil {
		log.Fatalf("failed to create db instance: %v", err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Print("success migrate db")
}
