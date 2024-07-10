package usersRoutes

import (
	"github.com/labstack/echo/v4"
	usersHandler "github.com/malikilamalik/fremont/internal/deliveries/http/v1/modules/users/handler"
)

func Init(app *echo.Group) {
	var (
		users   = app.Group("/users")
		handler = usersHandler.Init()
	)

	users.POST("/login", handler.Login)
	users.POST("/register", handler.Login)
}
