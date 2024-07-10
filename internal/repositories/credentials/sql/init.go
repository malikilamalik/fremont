package credentialsRepository

import (
	"context"

	"github.com/jmoiron/sqlx"
	credentialsModel "github.com/malikilamalik/fremont/internal/repositories/credentials/model"
)

type credentialRepository struct {
	db *sqlx.DB
}

type Credential interface {
	Create(ctx context.Context, data credentialsModel.Credential) error
	GetByUserID(ctx context.Context, userID string) (*credentialsModel.Credential, error)
}

func New(db *sqlx.DB) Credential {
	return &credentialRepository{
		db: db,
	}
}
