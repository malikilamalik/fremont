package entity

import (
	"time"

	"github.com/lib/pq"
)

type Users struct {
	ID        string      `db:"id"`
	Name      string      `db:"name"`
	Password  string      `db:"password"`
	Tokens    *Tokens     `db:"tokens"`
	CreatedAt time.Time   `db:"created_at"`
	DeletedAt pq.NullTime `db:"deleted_at"`
}

func (t Users) Table() string {
	return "users"
}
