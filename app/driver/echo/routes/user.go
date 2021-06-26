package routes

import (
	"github.com/labstack/echo"
	"github.com/yudwig/echo-sample/app/adapter/controller"
	"net/http"
)

func initUserRoutes(g *echo.Group) error {
	c, err := controller.NewUserController()
	if err != nil {
		return err
	}

	// Get Users API
	g.GET("/", func(ctx echo.Context) error {
		res := c.GetUsers()
		if res.Error.Code == 0 {
			return ctx.JSON(http.StatusOK, res)
		}
		return ctx.JSON(http.StatusInternalServerError, res)
	})

	// Create User API
	g.POST("/", func(ctx echo.Context) error {
		res := c.CreateUser(ctx.FormValue("name"))
		if res.Error.Code == 0 {
			return ctx.JSON(http.StatusCreated, res)
		}
		return ctx.JSON(http.StatusInternalServerError, res)
	})

	// Update User Name API
	g.PATCH("/:id", func(ctx echo.Context) error {
		res := c.UpdateUserName(ctx.Param("id"), ctx.FormValue("name"))
		if res.Error.Code == 0 {
			return ctx.JSON(http.StatusOK, res)
		}
		return ctx.JSON(http.StatusInternalServerError, res)
	})

	// Delete User Name API
	g.DELETE("/:id", func(ctx echo.Context) error {
		res := c.DeleteUser(ctx.Param("id"))
		if res.Error.Code == 0 {
			return ctx.JSON(http.StatusNoContent, res)
		}
		return ctx.JSON(http.StatusInternalServerError, res)
	})
	return nil
}
