package userapi

import "go.uber.org/fx"

var Module = fx.Module("api-user", fx.Invoke(
	Register,
	Login,
	Logout,
	CreateBank,
	GetUsers,
	GetProfile,
))
