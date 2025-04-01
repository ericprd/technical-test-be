package service

import (
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	usersvc.New,
))
