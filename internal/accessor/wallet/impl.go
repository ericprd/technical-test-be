package walletrepo

import (
	"github.com/ericprd/technical-test/database"
	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
)

type impl struct {
	db *database.DB
}

type Wallet interface {
	Register(spec walletdomain.Request) error
}

func New(db *database.DB) Wallet {
	return &impl{db}
}
