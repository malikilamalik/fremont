package postgresqlpkg

import (
	"fremont/config"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgreSQL() *sqlx.DB {
	config := config.PostgreSQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		Sslmode:  os.Getenv("SQL_SSL_MODE"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Database: os.Getenv("SQL_DATABASE"),
	}

	db, err := sqlx.Connect("postgres", config.FormatDSN())
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
