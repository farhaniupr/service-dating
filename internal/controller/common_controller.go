package controller

import (
	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/response"
	"github.com/labstack/echo/v4"
)

type CommonController struct {
	echoHandler library.RequestHandler
	env         library.Env
}

func ModuleCommonController(echoHandler library.RequestHandler,
	env library.Env) CommonController {
	return CommonController{
		echoHandler: echoHandler,
		env:         env,
	}
}

func (e CommonController) HealthCheck(c echo.Context) error {
	return response.ResponseInterface(c, 200, "OK", "Health Check")
}
