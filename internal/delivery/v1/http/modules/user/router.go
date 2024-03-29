package userroute

import (
	userhandler "fremont/internal/delivery/v1/http/modules/user/handler"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	var (
		h = userhandler.New()
	)
	publicRoute := g.Group("/users")
	publicRoute.GET("", h.Login)
}
