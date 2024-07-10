package usersRepository

import (
	"context"

	usersModel "github.com/malikilamalik/fremont/internal/repositories/users/model"
)

func (ur userRepository) Update(ctx context.Context, data usersModel.User) error {
	return nil
}
