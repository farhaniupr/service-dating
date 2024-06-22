package library

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestHandler struct {
	Echo *echo.Echo
}

func ModuleEcho() RequestHandler {
	// initial
	engine := echo.New()

	// middleware
	engine.Use(middleware.RequestID())

	return RequestHandler{Echo: engine}
}
