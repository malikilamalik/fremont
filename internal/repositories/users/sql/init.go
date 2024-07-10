package usersRepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	usersModel "github.com/malikilamalik/fremont/internal/repositories/users/model"
)

type userRepository struct {
	db *sqlx.DB
}

type User interface {
	Create(ctx context.Context, data usersModel.User) error
	Update(ctx context.Context, data usersModel.User) error
	GetByEmail(ctx context.Context, email string) (*usersModel.User, error)
}

func New(db *sqlx.DB) User {
	return &userRepository{
		db: db,
	}
}
