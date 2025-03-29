package main

import (
	"context"
	"log"

	"github.com/ericprd/technical-test/config"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file, continuing without it")
	}
}

func run(
	ls fx.Lifecycle,
	router chi.Router,
) {
	ls.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Print("application started")

			config.NewRouter(router)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Print("application stopped")
			return nil
		},
	})
}

func main() {
	app := fx.New(
		config.Module,
		fx.Invoke(run),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal("Error starting app:", err)
	}

	// Wait for the app to stop gracefully
	if err := app.Stop(context.Background()); err != nil {
		log.Fatal("Error stopping app:", err)
	}
}
