package userrepo

import (
	"github.com/ericprd/technical-test/database"
	"github.com/ericprd/technical-test/internal/domain/filter"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

type impl struct {
	db *database.DB
}

type User interface {
	Register(spec userdomain.RegisterSpec) error
	Login(username string) (userdomain.Profile, error)
	GetAllUser(f filter.FilterSpec) ([]userdomain.Profile, error)
	GetUser(id string) (*userdomain.Profile, error)
}

func New(db *database.DB) User {
	return &impl{db}
}
