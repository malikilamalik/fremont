package usersService

import (
	usersRepository "fremont/internal/repository/users"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	userRepo usersRepository.UserRepository
}

type UserService interface {
}

func NewUserService(db *sqlx.DB) UserService {
	return &userService{
		userRepo: usersRepository.NewRepository(db),
	}
}
