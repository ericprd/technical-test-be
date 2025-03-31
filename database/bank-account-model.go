package database

import "time"

type BankAccount struct {
	ID            string    `gorm:"primaryKey; not null"`
	UserID        string    `gorm:"uniqueIndex"`
	BankName      string    `gorm:"size:100"`
	AccountNumber string    `gorm:"size:50;unique"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
