package userdomain

import "time"

type Profile struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     *string   `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
