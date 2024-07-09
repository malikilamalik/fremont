package usersRepository

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        int64       `db:"id"`
	Name      string      `db:"name"`
	Email     string      `db:"email"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
}

func (u User) TableName() string {
	return `users`
}
