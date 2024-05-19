package entity

import (
	"time"
)

type Tokens struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Token     string    `db:"token"`
	ExpiredAt time.Time `db:"expired_at"`
}

func (t Tokens) Table() string {
	return "tokens"
}
