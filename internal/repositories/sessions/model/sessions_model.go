package sessionsModel

import "time"

type Session struct {
	ID        string    `db:"id"`
	Token     string    `db:"token"`
	ExpiredAt time.Time `db:"expired_at"`
	IPAddress string    `db:"ip_address"`
	UserAgent string    `db:"user_agent"`
	UserID    string    `db:"user_id"`
}

func (u Session) TableName() string {
	return `sessions`
}
