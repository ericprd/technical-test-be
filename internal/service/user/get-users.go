package usersvc

import (
	"fmt"

	"github.com/ericprd/technical-test/internal/domain/filter"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) GetAllUser(f filter.FilterSpec) ([]userdomain.Profile, error) {
	users, err := i.userRepo.GetAllUser(f)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}

	return users, nil
}
