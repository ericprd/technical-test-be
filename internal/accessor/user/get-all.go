package userrepo

import (
	"github.com/ericprd/technical-test/database"
	"github.com/ericprd/technical-test/internal/domain/filter"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

func (i *impl) GetAllUser(f filter.FilterSpec) ([]userdomain.Profile, error) {
	var users []userdomain.Profile

	q := i.db.
		Model(&database.User{}).
		Preload("Wallet").
		Preload("Bank").
		Omit("Password")

	if f.Username != "" {
		q = q.Where("users.username ILIKE ?", "%"+f.Username+"%")
	}

	if f.AccountName != "" || f.AccountNumber != "" {
		q = q.Joins("LEFT JOIN bank_accounts ON bank_accounts.user_id = users.id")

		if f.AccountName != "" {
			q = q.Where("bank_accounts.bank_name ILIKE ?", "%"+f.AccountName+"%")
		}

		if f.AccountNumber != "" {
			q = q.Where("bank_accounts.account_number = ?", f.AccountNumber)
		}
	}

	if !f.StartDate.IsZero() && !f.EndDate.IsZero() {
		q = q.Where("created_at BETWEEN ? AND ?", f.StartDate, f.EndDate)
	} else if !f.StartDate.IsZero() {
		q = q.Where("created_at >= ?", f.StartDate)
	} else if !f.EndDate.IsZero() {
		q = q.Where("created_at <= ?", f.EndDate)
	}

	if err := q.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
