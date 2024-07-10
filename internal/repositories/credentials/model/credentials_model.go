package credentialsModel

import (
	"time"

	"github.com/lib/pq"
)

type Credential struct {
	ID        string      `db:"id"`
	Active    string      `db:"active"`
	Hash      time.Time   `db:"hash"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
	UserID    string      `db:"user_id"`
}

func (u Credential) TableName() string {
	return `credentials`
}
