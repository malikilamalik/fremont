package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	usersController "github.com/malikilamalik/freemont/internal/delivery/http/v1/controller/users"
)

func Init(app *echo.Echo, db *sqlx.DB, val *validator.Validate) {
	v1 := app.Group("/v1")
	usersController.Init(v1, db, val)
}
