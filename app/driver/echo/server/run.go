package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yudwig/echo-sample/app/driver/echo/routes"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Logger().Printf("%#v", c.Request())
			err := next(c)
			c.Logger().Printf("%#v", c.Response())
			return err
		}
	})
	err := routes.Init(e)
	if err != nil {
		e.Logger.Fatal("failed to init router")
		e.Logger.Fatal(err)
		return
	}
	e.Logger.Fatal(e.Start("localhost:8000"))
}
