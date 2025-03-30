package config

import "go.uber.org/fx"

var Module = fx.Module(
	"config",
	fx.Provide(
		initRouter,
		newRedis,
	),
	fx.Invoke(InitConfig),
)
