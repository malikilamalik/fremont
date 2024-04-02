package usersql

import (
	"context"

	model "fremont/internal/entity"
	"fremont/pkg/common"
)

func (userRepository) GetList(ctx context.Context, param common.QueryParameter) ([]*model.User, error) {
	return nil, nil
}
