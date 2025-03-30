package main

import (
	"context"
	"log"

	"github.com/ericprd/technical-test/config"
	"github.com/ericprd/technical-test/database"
	"github.com/ericprd/technical-test/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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
	rdb *redis.Client,
) {
	ls.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Print("application started")

			_, err := rdb.Ping(ctx).Result()
			if err != nil {
				log.Fatalf("Could not connect to Redis: %v", err)
			}

			config.NewRouter(router)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Print("application stopped")
			rdb.Close()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		config.Module,
		database.Module,
		api.Module,
		fx.Invoke(run),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal("Error starting app:", err)
	}

	if err := app.Stop(context.Background()); err != nil {
		log.Fatal("Error stopping app:", err)
	}
}
