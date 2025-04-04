package database

import "time"

type BankAccount struct {
	ID            string    `gorm:"primaryKey; not null"`
	UserID        string    `gorm:"uniqueIndex"`
	AccountName   string    `gorm:"size:100"`
	AccountNumber string    `gorm:"size:50;unique"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
