package walletrepo

import walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"

func (i *impl) Register(spec walletdomain.Wallet) error {
	return i.db.Table("wallets").Omit("UpdatedAt").Create(&spec).Error
}
