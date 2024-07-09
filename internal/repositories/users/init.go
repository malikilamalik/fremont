package usersRepository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	Create(ctx context.Context, data User) error
	Update(ctx context.Context, data User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

func New(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
