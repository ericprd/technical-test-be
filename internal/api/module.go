package api

import (
	userapi "github.com/ericprd/technical-test/internal/api/user"
	walletapi "github.com/ericprd/technical-test/internal/api/wallet"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	fx.Provide(
		validator.New,
	),
	userapi.Module,
	walletapi.Module,
)
