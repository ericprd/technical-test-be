package walletdomain

type UpdateSpec struct {
	Amount int64 `json:"amount" validate:"required"`
	UserID string
}
