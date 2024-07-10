package sessionsRepository

import (
	"context"
	"fmt"
	"time"

	sessionsModel "github.com/malikilamalik/fremont/internal/repositories/sessions/model"
)

func (ur sessionRepository) GetByUserID(ctx context.Context, userID string) (*sessionsModel.Session, error) {
	var (
		resp sessionsModel.Session
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
