package postgresql

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/malikilamalik/freemont/config"
)

func New() *sqlx.DB {
	config := config.PostgreSQLConfig{
		Host:     os.Getenv("POSTGRE_DB_HOST"),
		Param:    os.Getenv("POSTGRE_DB_PARAMS"),
		User:     os.Getenv("POSTGRE_DB_USERNAME"),
		Password: os.Getenv("POSTGRE_DB_PASSWORD"),
		Port:     os.Getenv("POSTGRE_DB_PORT"),
		Database: os.Getenv("POSTGRE_DB_NAME"),
	}
	dsn := config.FormatDSN()
	db, err := sqlx.Open("pgx", dsn)

	if err != nil {
		log.Println("m=GetPool,msg=connection has failed", err)
	}

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return db
}
