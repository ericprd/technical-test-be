package userdomain

import (
	bankdomain "github.com/ericprd/technical-test/internal/domain/bank"
	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
)

type Profile struct {
	ID       string  `json:"id"`
	FullName string  `json:"full_name"`
	Username string  `json:"username"`
	Email    *string `json:"email"`
	Password string  `json:"password,omitempty"`

	Wallet *walletdomain.Wallet    `json:"wallet" gorm:"foreignKey:UserID"`
	Bank   *bankdomain.BankAccount `json:"bank" gorm:"foreignKey:UserID"`
}
