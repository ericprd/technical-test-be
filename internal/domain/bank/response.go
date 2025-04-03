package bankdomain

type BankAccount struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
}
