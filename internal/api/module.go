package api

import (
	helloworld "github.com/ericprd/technical-test/internal/api/hello-world"
	userapi "github.com/ericprd/technical-test/internal/api/user"
	walletapi "github.com/ericprd/technical-test/internal/api/wallet"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	fx.Provide(
		validator.New,
	),
	helloworld.Module,
	userapi.Module,
	walletapi.Module,
)
