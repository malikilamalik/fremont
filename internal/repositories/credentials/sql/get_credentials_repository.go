package credentialsRepository

import (
	"context"
	"fmt"
	"time"

	credentialsModel "github.com/malikilamalik/fremont/internal/repositories/credentials/model"
)

func (ur credentialRepository) GetByUserID(ctx context.Context, userID string) (*credentialsModel.Credential, error) {
	var (
		resp credentialsModel.Credential
		err  error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "user_id" = $1`, resp.TableName())
	err = ur.db.GetContext(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
