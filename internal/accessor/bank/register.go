package bankrepo

import (
	bankdomain "github.com/ericprd/technical-test/internal/domain/bank"
)

func (i *impl) Register(spec bankdomain.RegisterSpec) error {
	if err := i.db.Table("bank_accounts").Omit("UpdatedAt").Create(&spec).Error; err != nil {
		return err
	}

	return nil
}
