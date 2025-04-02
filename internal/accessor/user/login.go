package userrepo

import (
	"github.com/ericprd/technical-test/database"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) Login(username string) (userdomain.Profile, error) {
	var user userdomain.Profile
	if err := i.db.Model(&database.User{}).First(&user, "username=?", username).Error; err != nil {
		return user, err
	}

	return user, nil
}
