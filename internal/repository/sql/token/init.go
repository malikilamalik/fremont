package tokensql

import (
	"context"
	"database/sql"
	"fremont/internal/repository/model"
)

type tokenRepository struct {
	db *sql.DB
}

type TokenRepository interface {
	Get(ctx context.Context) (*model.Token, error)
}

func NewTokenRepository(db *sql.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}
