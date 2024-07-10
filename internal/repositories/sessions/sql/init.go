package sessionsRepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	sessionsModel "github.com/malikilamalik/fremont/internal/repositories/sessions/model"
)

type sessionRepository struct {
	db *sqlx.DB
}

type Session interface {
	Create(ctx context.Context, data sessionsModel.Session) error
	GetByUserID(ctx context.Context, userID string) (*sessionsModel.Session, error)
}

func New(db *sqlx.DB) Session {
	return &sessionRepository{
		db: db,
	}
}
