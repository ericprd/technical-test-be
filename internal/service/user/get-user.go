package usersvc

import (
	"fmt"

	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) GetUser(id string) (*userdomain.Profile, error) {
	user, err := i.userRepo.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return user, nil
}
