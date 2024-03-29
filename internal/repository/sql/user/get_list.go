package usersql

import (
	"context"
	"fremont/internal/repository/model"
	"fremont/pkg/common"
)

func (userRepository) GetList(ctx context.Context, param common.QueryParameter) ([]*model.User, error) {
	return nil, nil
}
