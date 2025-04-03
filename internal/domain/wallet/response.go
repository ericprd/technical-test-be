package walletdomain

type Wallet struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Balance int64  `json:"balance"`
}
