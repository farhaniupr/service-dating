package library

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ModuleConfig),
	fx.Provide(ModuleEcho),
	fx.Provide(ModuleDatabase),
	fx.Provide(ModuleLogger),
	fx.Provide(ModuleRedis),
)
