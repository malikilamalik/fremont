package usersController

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init(group *echo.Group, val *validator.Validate) {
	user := group.Group("/user")

	//Private
	publicRouter := user
	publicRouter.POST("/register", func(c echo.Context) error {
		return nil
	})
	publicRouter.POST("/login", func(c echo.Context) error {
		return nil
	})
}
