package helloworld

import "go.uber.org/fx"

var Module = fx.Module("hello", fx.Invoke(
	Get,
))
