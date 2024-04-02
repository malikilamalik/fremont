package model

import "time"

type Token struct {
	ID        string    `db:"id"`
	Token     string    `db:"string"`
	Expires   int64     `db:"expires"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (g Token) TableName() string {
	return `token`
}
