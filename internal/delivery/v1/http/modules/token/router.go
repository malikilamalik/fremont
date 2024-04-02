package tokenrouter

import (
	tokenhandler "fremont/internal/delivery/v1/http/modules/token/handler"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	var (
		h = tokenhandler.New()
	)
	publicRoute := g.Group("/auth")
	publicRoute.POST("/login", h.Login)
	publicRoute.POST("/logout", h.Login)
	publicRoute.POST("/refresh", h.Refresh)
}
