package accessor

import (
	redisrepo "github.com/ericprd/technical-test/internal/accessor/redis"
	userrepo "github.com/ericprd/technical-test/internal/accessor/user"
	"go.uber.org/fx"
)

var Module = fx.Module("accessor", fx.Provide(
	redisrepo.New,
	userrepo.New,
))
