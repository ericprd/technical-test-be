package userrepo

import (
	"github.com/ericprd/technical-test/database"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) GetAllUser() ([]userdomain.Profile, error) {
	var users []userdomain.Profile

	if err := i.db.Model(&database.User{}).Omit("Password").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
