package usersvc

import (
	"fmt"

	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) GetAllUser() ([]userdomain.Profile, error) {
	users, err := i.userRepo.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}
