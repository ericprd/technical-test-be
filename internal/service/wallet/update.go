package walletsvc

import (
	"fmt"

	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
	timeutil "github.com/ericprd/technical-test/internal/utils/time"
)

func (i *impl) Update(spec walletdomain.UpdateSpec) error {
	wallet, err := i.walletRepo.GetWallet(spec.UserID)
	if err != nil {
		return fmt.Errorf("failed to get wallet: %v", err)
	}

	wallet.Balance += spec.Amount
	wallet.UpdatedAt = timeutil.UTCNow()

	if err := i.walletRepo.Update(wallet); err != nil {
		return fmt.Errorf("failed to update wallet: %v", err)
	}

	return nil
}
