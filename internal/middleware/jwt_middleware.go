package middleware

import (
	"strings"

	"github.com/farhaniupr/dating-api/internal/service"
	"github.com/farhaniupr/dating-api/package/library"
	"github.com/labstack/echo/v4"
)

// JWTAuthMiddleware middleware for jwt authentication
type JWTAuthMiddleware struct {
	service service.JWTAuthMethodService
	env     library.Env
}

// ModuleJWTAuthMiddleware creates new jwt auth middleware
func ModuleJWTAuthMiddleware(
	env library.Env,
	service service.JWTAuthMethodService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: service,
		env:     env,
	}
}

// Setup sets up jwt auth middleware
func (m JWTAuthMiddleware) Setup() {}

// Handler handles middleware functionality
func (m JWTAuthMiddleware) Handler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				user, authorized, err := m.service.Authorize(authToken)
				if err != nil {
					library.Writelog(c, m.env, "err", err.Error())
					return nil
				}

				c.Set("data_jwt", user)
				_ = user
				if authorized {
					next(c)
					return nil
				}
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return nil
			}
			c.JSON(401, map[string]interface{}{
				"error": "you are not authorized",
			})
			return nil
		}
	}
}
