package bankdomain

import "time"

type RegisterSpec struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	BankName      string    `json:"bank_name" validate:"required"`
	AccountNumber string    `json:"account_number" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at"`
}
