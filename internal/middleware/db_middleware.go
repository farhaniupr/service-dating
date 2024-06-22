package middleware

import (
	"log"
	"net/http"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/constants"
	"github.com/labstack/echo/v4"
)

// DatabaseTrx middleware for transactions support for database
type DatabaseTrx struct {
	handler library.RequestHandler
	db      library.Database
	env     library.Env
}

// statusInList function checks if context writer status is in provided list
func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// ModuleDatabase creates new database transactions middleware
func ModuleDatabase(
	handler library.RequestHandler,
	db library.Database,
	env library.Env,
) DatabaseTrx {
	return DatabaseTrx{
		handler: handler,
		db:      db,
		env:     env,
	}
}

func (m DatabaseTrx) HandlerDB() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println("beginning database transaction")

			txHandle := m.db.MysqlDB.WithContext(c.Request().Context()).Begin()

			defer func() {
				if r := recover(); r != nil {
					txHandle.Rollback()
					log.Println("rollback database transaction")
					return
				}
			}()

			c.Set(constants.DBTransaction, txHandle)
			if err := next(c); err != nil {
				log.Println("commit err : ", err.Error())
				return err
			}

			// commit transaction on success status
			if statusInList(c.Response().Status,
				[]int{http.StatusOK, http.StatusCreated, http.StatusNoContent, http.StatusPermanentRedirect, http.StatusTemporaryRedirect}) {

				log.Println("commit database transaction")
				if err := txHandle.Commit().Error; err != nil {
					log.Println("commit err : ", err.Error())
					return err
				}

			} else {

				log.Println("rolling back database transaction")
				txHandle.Rollback()

				return nil

			}

			return nil
		}
	}
}

func (m DatabaseTrx) HandlerDBContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			c.Set("ctx", c.Request().Context())

			if err := next(c); err != nil {
				log.Println("err get contenxt : ", err.Error())
				return err
			}

			return nil
		}
	}
}
