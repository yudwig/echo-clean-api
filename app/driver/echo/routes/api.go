package routes

import (
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) error {
	err := initUserRoutes(e.Group("/user"))
	if err != nil {
		return err
	}
	return nil
}
