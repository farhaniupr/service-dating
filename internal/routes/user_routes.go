package routes

import (
	"github.com/farhaniupr/dating-api/internal/controller"
	"github.com/farhaniupr/dating-api/internal/middleware"
	"github.com/farhaniupr/dating-api/package/library"
)

// UserRoutes struct routes
type UserRoutes struct {
	handler        library.RequestHandler
	userController controller.UserController

	middlewareDB  middleware.DatabaseTrx
	middlewareJwt middleware.JWTAuthMiddleware
}

// Setup routes
func (c UserRoutes) Setup() {
	api := c.handler.Echo.Group("user/")
	{

		api.POST("login", c.userController.Login)
		api.POST("register", c.userController.Register, c.middlewareDB.HandlerDB())
		api.GET("detail/:phone", c.userController.DetailUser, c.middlewareDB.HandlerDBContext(), c.middlewareJwt.Handler())
		api.PATCH("update/:phone", c.userController.UpdateUser, c.middlewareDB.HandlerDB(), c.middlewareJwt.Handler())

		api.GET("find-date", c.userController.Finddate, c.middlewareJwt.Handler(), c.middlewareJwt.Handler())
		api.GET("swift-right/:phone_target", c.userController.SwiftRight, c.middlewareDB.HandlerDB(), c.middlewareJwt.Handler())
		api.GET("swift-left", c.userController.Finddate, c.middlewareJwt.Handler(), c.middlewareJwt.Handler())

		api.POST("buy-premium", c.userController.BuyPremium, c.middlewareDB.HandlerDB(), c.middlewareJwt.Handler())

	}
}

// ModuleUserRoutes
func ModuleUserRoutes(
	handler library.RequestHandler,
	userController controller.UserController,
	middlewareDB middleware.DatabaseTrx,
	middlewareJwt middleware.JWTAuthMiddleware,

) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
		middlewareDB:   middlewareDB,
		middlewareJwt:  middlewareJwt,
	}
}
