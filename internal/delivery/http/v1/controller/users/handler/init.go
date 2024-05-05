package userHandler

import (
	usersService "fremont/internal/service/users"

	"github.com/jmoiron/sqlx"
)

type userHandler struct {
	userService usersService.UserService
}

func NewHandler(db *sqlx.DB) *userHandler {
	return &userHandler{
		userService: usersService.NewUserService(db),
	}
}
