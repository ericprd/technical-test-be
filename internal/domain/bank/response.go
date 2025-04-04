package bankdomain

type BankAccount struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}
