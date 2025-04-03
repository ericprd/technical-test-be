package walletrepo

import walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"

func (i *impl) Update(spec walletdomain.Wallet) error {
	if err := i.db.Table("wallets").
		Where("id=?", spec.ID).
		Updates(&spec).
		Error; err != nil {
		return err
	}

	return nil
}
