package walletdomain

import "time"

type Request struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   int64     `json:"balance"`
	UpdatedAt time.Time `json:"updated_at"`
}
