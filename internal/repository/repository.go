package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ModuleUserRepository),
	fx.Provide(ModuleUserLiked),
)
