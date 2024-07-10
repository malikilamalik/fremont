package usersModel

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        string      `db:"id"`
	Name      string      `db:"name"`
	Email     string      `db:"email"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
}

func (u User) TableName() string {
	return `users`
}
