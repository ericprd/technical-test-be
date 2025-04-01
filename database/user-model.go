package database

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;not null"`
	FullName  string    `gorm:"size:100"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;default=null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
