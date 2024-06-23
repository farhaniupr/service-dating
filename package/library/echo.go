package library

import (
	"time"

	"github.com/farhaniupr/dating-api/resource/constants"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type RequestHandler struct {
	Echo *echo.Echo
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func dateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(constants.LayoutDate, fl.Field().String())
	return err == nil
}

func ModuleEcho() RequestHandler {
	// initial
	engine := echo.New()

	// middleware
	engine.Use(middleware.RequestID())

	validator := validator.New()
	err := validator.RegisterValidation("date", dateValidation)
	if err != nil {
		log.Fatal(err.Error())
	}

	engine.Validator = &CustomValidator{Validator: validator}

	return RequestHandler{Echo: engine}
}
