package walletrepo

import walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"

func (i *impl) GetWallet(userID string) (walletdomain.Wallet, error) {
	var wallet walletdomain.Wallet

	if err := i.db.Table("wallets").Omit("UpdatedAt").First(&wallet, "user_id=?", userID).Error; err != nil {
		return wallet, err
	}

	return wallet, nil
}
