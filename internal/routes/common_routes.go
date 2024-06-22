package routes

import (
	"github.com/farhaniupr/dating-api/internal/controller"
	"github.com/farhaniupr/dating-api/package/library"
)

// CommonRoutes struct routes
type CommonRoutes struct {
	handler          library.RequestHandler
	commonController controller.CommonController
}

// Setup routes
func (c CommonRoutes) Setup() {
	api := c.handler.Echo.Group("/")
	{
		api.GET("health-check", c.commonController.HealthCheck)
	}
}

// ModuleCommonRoutes
func ModuleCommonRoutes(
	handler library.RequestHandler,
	commonController controller.CommonController,
) CommonRoutes {
	return CommonRoutes{
		handler:          handler,
		commonController: commonController,
	}
}
