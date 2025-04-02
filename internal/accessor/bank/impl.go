package bankrepo

import (
	"github.com/ericprd/technical-test/database"
	bankdomain "github.com/ericprd/technical-test/internal/domain/bank"
)

type impl struct {
	db *database.DB
}

type Bank interface {
	Register(spec bankdomain.RegisterSpec) error
}

func New(db *database.DB) Bank {
	return &impl{db}
}
