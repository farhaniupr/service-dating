package controller

import (
	"context"
	"net/http"

	"github.com/farhaniupr/dating-api/internal/helper"
	"github.com/farhaniupr/dating-api/internal/service"
	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/constants"
	"github.com/farhaniupr/dating-api/resource/model"
	"github.com/farhaniupr/dating-api/resource/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	env          library.Env
	userService  service.UserMethodService
	commonHelper helper.CommonHelper
}

func ModuleUserController(
	echoHandler library.RequestHandler,
	env library.Env,
	userService service.UserMethodService,
	commonHelper helper.CommonHelper,
) UserController {
	return UserController{
		env:          env,
		userService:  userService,
		commonHelper: commonHelper,
	}
}

func (u UserController) DetailUser(c echo.Context) error {

	id := c.Param("phone")

	result, err := u.userService.DetailUser(c.Get("ctx").(context.Context), id)
	if err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return response.ResponseInterface(c, 500, err.Error(), constants.InternalServerError)
	}

	return response.ResponseInterface(c, 200, result, "Detail Userr")
}

func (u UserController) ListUser(c echo.Context) error {
	var pagenation model.Pagenation

	if err := c.Bind(&pagenation); err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return response.ResponseInterfaceTotal(c, 500, err.Error(), constants.BadRequest, 0)
	}

	result, total, err := u.userService.ListUser(c.Get("ctx").(context.Context), pagenation.Page, pagenation.Limit)
	if err != nil {
		return response.ResponseInterfaceTotal(c, 500, err.Error(), constants.InternalServerError, 0)

	}

	return response.ResponseInterfaceTotal(c, 200, result, "List User", int(total))
}

func (u UserController) UpdateUser(c echo.Context) error {
	var dataReq model.User

	id := c.Param("phone")

	if err := c.Bind(&dataReq); err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return response.ResponseInterface(c, 500, err.Error(), constants.BadRequest)
	}

	if err := c.Validate(dataReq); err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return response.ResponseInterfaceError(c, http.StatusBadRequest, library.GetValueBetween(err.Error(), "Error:", "tag"), constants.BadRequest)
	}

	tx := c.Get(constants.DBTransaction).(*gorm.DB)

	result, err := u.userService.UpdateUser(tx, dataReq, id)
	if err != nil {
		return response.ResponseInterface(c, 500, err.Error(), constants.InternalServerError)
	}

	return response.ResponseInterface(c, 200, result, "Update User")
}

func (u UserController) Login(c echo.Context) error {

	var dataReq model.User

	err := c.Bind(&dataReq)
	if err != nil {
		return response.ResponseInterface(c, http.StatusBadRequest, err.Error(), "Bad Request")
	}

	_, resultUser, err := u.userService.Login(c.Request().Context(), dataReq)
	if err != nil {
		switch err.Error() {
		case "account not found":
			return response.ResponseInterface(c, 200, err.Error(), err.Error())
		case "password is wrong":
			return response.ResponseInterface(c, 200, err.Error(), err.Error())
		default:
			return response.ResponseInterface(c, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		}
	}

	return response.ResponseInterface(c, 200, resultUser, "Login Success")
}

func (u UserController) Register(c echo.Context) error {

	var dataReq model.User

	err := c.Bind(&dataReq)
	if err != nil {
		return response.ResponseInterface(c, http.StatusBadRequest, err.Error(), "Bad Request")
	}

	if err := c.Validate(dataReq); err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return response.ResponseInterfaceError(c, http.StatusBadRequest, library.GetValueBetween(err.Error(), "Error:", "tag"), constants.BadRequest)
	}

	resultUser, err := u.userService.StoreUser(c.Get(constants.DBTransaction).(*gorm.DB), dataReq)
	if err != nil {

		return response.ResponseInterface(c, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	return response.ResponseInterface(c, 201, resultUser, "Register Success")
}

func (u UserController) SwiftRight(c echo.Context) error {

	var dataReq model.User

	err := c.Bind(&dataReq)
	if err != nil {
		return response.ResponseInterface(c, http.StatusBadRequest, err.Error(), "Bad Request")
	}

	resultUser, err := u.userService.StoreUser(c.Get(constants.DBTransaction).(*gorm.DB), dataReq)
	if err != nil {

		return response.ResponseInterface(c, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	return response.ResponseInterface(c, 200, resultUser, "Register Success")
}

func (u UserController) SwiftLeft(c echo.Context) error {

	var dataReq model.User

	err := c.Bind(&dataReq)
	if err != nil {
		return response.ResponseInterface(c, http.StatusBadRequest, err.Error(), "Bad Request")
	}

	resultUser, err := u.userService.StoreUser(c.Get(constants.DBTransaction).(*gorm.DB), dataReq)
	if err != nil {

		return response.ResponseInterface(c, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	return response.ResponseInterface(c, 200, resultUser, "Register Success")
}

func (u UserController) Finddate(c echo.Context) error {

	resultUser, err := u.userService.FindDate(c.Request().Context(), c.Get("data_jwt").(map[string]interface{}))
	if err != nil {
		return response.ResponseInterface(c, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	return response.ResponseInterface(c, 200, resultUser, "Data UserDate")
}
