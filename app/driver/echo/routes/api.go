package routes

import (
	"github.com/labstack/echo"
	"github.com/yudwig/echo-sample/app/adapter/controller"
	"net/http"
)

func initUserRoutes(g *echo.Group) {
	c := controller.NewUserController()
	g.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, c.CreateUser("my name"))
	})
}

func Init(e *echo.Echo) {
	initUserRoutes(e.Group("/user"))
}
