package service

import (
	authsvc "github.com/ericprd/technical-test/internal/service/auth"
	"github.com/ericprd/technical-test/internal/service/middlewaresvc"
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	usersvc.New,
	authsvc.New,
	middlewaresvc.New,
))
