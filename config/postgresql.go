package config

import "fmt"

type PostgreSQLConfig struct {
	Host     string
	Sslmode  string
	User     string
	Password string
	Database string
	Port     string
}

func (m *PostgreSQLConfig) FormatDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", m.User, m.Password, m.Host, m.Port, m.Database, m.Sslmode)
}
