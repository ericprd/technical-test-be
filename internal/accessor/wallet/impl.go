package walletrepo

import (
	"github.com/ericprd/technical-test/database"
	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
)

type impl struct {
	db *database.DB
}

type Wallet interface {
	Register(spec walletdomain.Wallet) error
	GetWallet(userID string) (walletdomain.Wallet, error)
	Update(spec walletdomain.Wallet) error
}

func New(db *database.DB) Wallet {
	return &impl{db}
}
