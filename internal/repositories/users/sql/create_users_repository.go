package usersRepository

import (
	"context"
	"fmt"
	"time"

	usersModel "github.com/malikilamalik/fremont/internal/repositories/users/model"
)

func (ur userRepository) Create(ctx context.Context, data usersModel.User) error {
	var (
		err error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf(`INSERT INTO %s(id, email, name, created_at) VALUES (:id, :email, :name, :created_at)`, data.TableName())
	_, err = ur.db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	return err
}
