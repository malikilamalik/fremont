package user

import (
	"github.com/jmoiron/sqlx"
)

type userService struct {
	db *sqlx.DB
}

type UserService interface {
}

func NewUserService(db *sqlx.DB) UserService {
	return &userService{
		db: db,
	}
}
