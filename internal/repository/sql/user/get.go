package usersql

import (
	"context"
	"fmt"
	"fremont/internal/repository/model"
	"fremont/pkg/common"
	"time"
)

func (ur userRepository) Get(ctx context.Context, param common.QueryParameter) (*model.User, error) {
	var (
		res model.User
	)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT id FROM public.user"

	err := ur.db.GetContext(ctx, &res, query)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer ur.db.Close()
	return &res, nil
}
