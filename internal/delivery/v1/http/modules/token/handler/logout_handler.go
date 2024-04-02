package tokenhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h tokenHandler) Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, "A")
}
