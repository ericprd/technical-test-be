package database

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;not null"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;default=null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `grom:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `grom:"default:CURRENT_TIMESTAMP"`
}
