package modulesV1

import (
	userroute "fremont/internal/delivery/v1/http/modules/user"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	g := e.Group("/v1")
	userroute.Init(g)
}
