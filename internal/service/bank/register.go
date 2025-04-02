package banksvc

import (
	"fmt"

	bankdomain "github.com/ericprd/technical-test/internal/domain/bank"
	"github.com/google/uuid"
)

func (i *impl) Register(spec bankdomain.RegisterSpec) error {
	spec.ID = uuid.NewString()

	if err := i.bankRepo.Register(spec); err != nil {
		return fmt.Errorf("failed to create user bank: %v", err)
	}

	return nil
}
