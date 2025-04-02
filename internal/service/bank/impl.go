package banksvc

import (
	bankrepo "github.com/ericprd/technical-test/internal/accessor/bank"
	bankdomain "github.com/ericprd/technical-test/internal/domain/bank"
)

type impl struct {
	bankRepo bankrepo.Bank
}

type Bank interface {
	Register(spec bankdomain.RegisterSpec) error
}

func New(
	bankRepo bankrepo.Bank,
) Bank {
	return &impl{bankRepo}
}
