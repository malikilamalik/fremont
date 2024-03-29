package usersql

import (
	"context"
	"fremont/internal/repository/model"
	"fremont/pkg/common"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	Get(ctx context.Context, param common.QueryParameter) (*model.User, error)
	GetList(ctx context.Context, param common.QueryParameter) ([]*model.User, error)
	Create(ctx context.Context, data model.User) error
	Update(ctx context.Context, data model.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
