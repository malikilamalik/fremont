package usersHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (userHandler) Register(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
