package database

import (
	"fmt"

	"github.com/ericprd/technical-test/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		config.POSTGRES_HOST, config.POSTGRES_USER, config.POSTGRES_PASS,
		config.POSTGRES_DB, config.POSTGRES_PORT,
	)

	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DB{db}, nil
}

func (db *DB) Migrate() error {
	return db.AutoMigrate(
		&User{},
		&Wallet{},
		&BankAccount{},
	)
}
