package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func test(c echo.Context) error {
	fmt.Println("test is called.")
	return c.JSON(http.StatusOK, c)
}

func main() {
	fmt.Println("start")
	e := echo.New()
	e.GET("/test", test)
	e.Start("localhost:8000")
}
