package sessionsRepository

import (
	"context"
	"fmt"
	"time"

	sessionsModel "github.com/malikilamalik/fremont/internal/repositories/sessions/model"
)

func (ur sessionRepository) Create(ctx context.Context, data sessionsModel.Session) error {
	var (
		err error
	)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf(`INSERT INTO %s(id, token, expired_at, ip_address, user_agent, user_id) VALUES (:id, :token, :expired_at, :ip_address, :user_agent, :user_id)`, data.TableName())
	_, err = ur.db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	return err
}
