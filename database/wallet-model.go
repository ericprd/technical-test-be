package database

import "time"

type Wallet struct {
	ID        string    `gorm:"primaryKey; not null"`
	UserID    string    `gorm:"uniqueIndex;not null"`
	Balance   float64   `gorm:"default:0"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
