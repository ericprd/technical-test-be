package database

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;not null"`
	FirstName string    `gorm:"size:100"`
	LastName  string    `gorm:"size:100"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;default=null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `grom:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `grom:"default:CURRENT_TIMESTAMP"`

	BankAccount BankAccount `gorm:"foreignKey:ID;references:UserID;constraint:OnDelete:CASCADE"`
	Wallet      Wallet      `gorm:"foreignKey:ID;references:UserID;constraint:OnDelete:CASCADE"`
}
