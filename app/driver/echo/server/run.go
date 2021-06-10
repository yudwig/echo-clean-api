package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yudwig/echo-sample/app/driver/echo/routes"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	routes.Init(e)
	e.Logger.Fatal(e.Start("localhost:8000"))
}
