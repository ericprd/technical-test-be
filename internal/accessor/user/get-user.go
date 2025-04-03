package userrepo

import (
	"github.com/ericprd/technical-test/database"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) GetUser(id string) (*userdomain.Profile, error) {
	var user *userdomain.Profile

	if err := i.db.
		Model(&database.User{}).
		Preload("Wallet").
		Preload("Bank").
		Omit("Password").
		First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}
