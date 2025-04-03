package walletsvc

import (
	walletrepo "github.com/ericprd/technical-test/internal/accessor/wallet"
	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
)

type impl struct {
	walletRepo walletrepo.Wallet
}

type Wallet interface {
	Update(spec walletdomain.UpdateSpec) error
}

func New(walletRepo walletrepo.Wallet) Wallet {
	return &impl{walletRepo}
}
