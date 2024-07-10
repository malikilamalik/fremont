package credentialsRepository

import (
	"context"
	"fmt"
	"time"

	credentialsModel "github.com/malikilamalik/fremont/internal/repositories/credentials/model"
)

func (ur credentialRepository) Create(ctx context.Context, data credentialsModel.Credential) error {
	var (
		err error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf(`INSERT INTO %s(id, active, hash, created_at, user_id) VALUES (:id, :active, :hash, :created_at, :user_id)`, data.TableName())
	_, err = ur.db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	return err
}
