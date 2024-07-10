package usersRepository

import (
	"context"
	"fmt"
	"time"

	usersModel "github.com/malikilamalik/fremont/internal/repositories/users/model"
)

func (ur userRepository) GetByEmail(ctx context.Context, email string) (*usersModel.User, error) {
	var (
		resp usersModel.User
		err  error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "email" = $1`, resp.TableName())
	err = ur.db.GetContext(ctx, &resp, query, email)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
