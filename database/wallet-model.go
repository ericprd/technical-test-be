package database

import "time"

type Wallet struct {
	ID        string    `gorm:"primaryKey; not null"`
	UserID    string    `gorm:"uniqueIndex"`
	Balance   float64   `gorm:"default:0"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
