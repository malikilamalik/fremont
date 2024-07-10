package usersHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (userHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
