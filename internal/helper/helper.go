package helper

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(ModuleCommonHelper),
	fx.Provide(ModuleEncryptHelper),
)
