package usersql

import (
	"context"
	model "fremont/internal/entity"
)

func (userRepository) Update(ctx context.Context, data model.User) error {
	return nil
}
