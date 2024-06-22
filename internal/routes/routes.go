package routes

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(ModuleCommonRoutes),
	fx.Provide(ModuleRoutes),
	fx.Provide(ModuleUserRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func ModuleRoutes(
	commonRoutes CommonRoutes,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		commonRoutes,
		userRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
