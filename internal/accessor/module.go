package accessor

import (
	bankrepo "github.com/ericprd/technical-test/internal/accessor/bank"
	redisrepo "github.com/ericprd/technical-test/internal/accessor/redis"
	userrepo "github.com/ericprd/technical-test/internal/accessor/user"
	walletrepo "github.com/ericprd/technical-test/internal/accessor/wallet"
	"go.uber.org/fx"
)

var Module = fx.Module("accessor", fx.Provide(
	redisrepo.New,
	userrepo.New,
	walletrepo.New,
	bankrepo.New,
))
