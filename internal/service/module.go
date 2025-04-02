package service

import (
	authsvc "github.com/ericprd/technical-test/internal/service/auth"
	banksvc "github.com/ericprd/technical-test/internal/service/bank"
	middlewaresvc "github.com/ericprd/technical-test/internal/service/middleware"
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	usersvc.New,
	authsvc.New,
	middlewaresvc.New,
	banksvc.New,
))
