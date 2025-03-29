package main

import (
	"log"

	"github.com/ericprd/technical-test/config"
	"github.com/ericprd/technical-test/database"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file, continuing without it")
	}
}

func main() {
	fx.New(
		config.Module,
		database.Module,
		fx.Invoke(func(
			db *database.DB,
		) {
			if err := db.Migrate(); err != nil {
				log.Fatalf("failed to migrate database: %s", err)
			}
		}),
	).Run()
}
