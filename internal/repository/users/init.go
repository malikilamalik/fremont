package usersRepository

import (
	"context"
	"fremont/internal/entity"
	"fremont/pkg/common"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	Get(ctx context.Context, param common.QueryParameter) (*entity.User, error)
	Create(ctx context.Context, data entity.User) error
}

func NewRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
