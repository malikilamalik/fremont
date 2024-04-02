package usersql

import (
	"context"
	model "fremont/internal/entity"
)

func (userRepository) Create(ctx context.Context, data model.User) error {
	return nil
}
