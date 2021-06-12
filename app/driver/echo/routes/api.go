package routes

import (
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	initUserRoutes(e.Group("/user"))
}
