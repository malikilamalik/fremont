package usersql

import (
	"context"
	"fremont/internal/repository/model"
	"fremont/pkg/common"
	"time"
)

func (ur userRepository) Get(ctx context.Context, param common.QueryParameter) (*model.User, error) {
	var (
		res   model.User
		query string
		err   error
	)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query = "SELECT id, email, first_name, last_name FROM public.user"

	err = ur.db.GetContext(ctx, &res, query)
	if err != nil {
		return nil, err
	}
	defer ur.db.Close()

	return &res, nil
}
