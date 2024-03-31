package userhandler

import (
	"context"
	usersql "fremont/internal/repository/sql/user"
	"fremont/pkg/common"
	postgresqlpkg "fremont/pkg/database/postgresql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h userHandler) Register(c echo.Context) error {
	db := postgresqlpkg.InitPostgreSQL()
	repository := usersql.NewUserRepository(db)
	a, _ := repository.Get(context.Background(), common.QueryParameter{})

	return c.JSON(http.StatusOK, a)
}
