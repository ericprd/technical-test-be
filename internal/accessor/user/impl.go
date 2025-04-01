package userrepo

import (
	"github.com/ericprd/technical-test/database"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

type impl struct {
	db *database.DB
}

type User interface {
	Register(spec userdomain.Request) error
}

func New(db *database.DB) User {
	return &impl{db}
}
