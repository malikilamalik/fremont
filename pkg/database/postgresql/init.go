package postgresql

import (
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/malikilamalik/fremont/config"
)

func New(config config.PostgreSQLConfig) *sqlx.DB {
	dsn := config.FormatDSN()
	db, err := sqlx.Open("pgx", dsn)

	if err != nil {
		log.Println("m=GetPool,msg=connection has failed", err)
	}

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	db.SetMaxOpenConns(70)
	db.SetMaxIdleConns(20)

	return db
}
