package usersql

import (
	"context"
	"fremont/internal/repository/model"
)

func (userRepository) Create(ctx context.Context, data model.User) error {
	return nil
}