package server

import (
	"github.com/labstack/echo/v4"
	usersRoutes "github.com/malikilamalik/fremont/internal/deliveries/http/v1/modules/users"
)

func Routes(app *echo.Echo) {
	var (
		v1 = app.Group("v1")
	)
	usersRoutes.Init(v1)
}
