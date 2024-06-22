package middleware

import "go.uber.org/fx"

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(ModuleDatabase),
	fx.Provide(ModuleMiddlewares),
	fx.Provide(ModuleJWTAuthMiddleware),
	fx.Provide(ModuleCorsMiddleware),
	fx.Provide(ModuleLogger),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// ModuleMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
func ModuleMiddlewares(
	jwtauthMiddleware JWTAuthMiddleware,
	corsMiddleware CorsMiddleware,
	// mongoMiddleware MongoMiddleware,
	// loggerMiddleware LoggerMiddlewre,
) Middlewares {
	return Middlewares{
		jwtauthMiddleware,
		corsMiddleware,
		// mongoMiddleware,
		// loggerMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
