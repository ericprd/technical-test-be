package userrepo

import userdomain "github.com/ericprd/technical-test/internal/domain/user"

func (i *impl) Register(spec userdomain.RegisterSpec) error {
	return i.db.Table("users").Create(&spec).Error
}
