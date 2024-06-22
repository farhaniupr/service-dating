package middleware

import (
	"net/http"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/labstack/echo/v4"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler library.RequestHandler
}

// ModuleCorsMiddleware creates new cors middleware
func ModuleCorsMiddleware(handler library.RequestHandler) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {

	m.handler.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			req := c.Request()
			res := c.Response()

			allowList := map[string]bool{
				"http://localhost:3000": true,
				"":                      true,
			}

			if origin := req.Header.Get("Origin"); allowList[origin] {
				res.Header().Add("Access-Control-Allow-Origin", origin)
				res.Header().Add("Access-Control-Allow-Methods", "*")
				res.Header().Add("Access-Control-Allow-Headers", "*")
				res.Header().Add("Content-Type", "application/json")

				if req.Method != "OPTIONS" {
					err := next(c)
					if err != nil {
						c.Error(err)
					}
				}
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Cors Origin")
			}

			return nil

		}
	})
}
