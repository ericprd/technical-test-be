package walletapi

import "go.uber.org/fx"

var Module = fx.Module("api-wallet", fx.Invoke(
	Update,
))
