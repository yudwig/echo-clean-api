package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	routes2 "github.com/yudwig/echo-sample/app/driver/echo/routes"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	routes2.Init(e)
	e.Logger.Fatal(e.Start("localhost:8000"))
}
