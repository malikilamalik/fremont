package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/malikilamalik/freemont/internal/delivery/http/v1/request"
)

type Users struct {
	ID        string        `db:"id"`
	Name      string        `db:"name"`
	Email     string        `db:"name"`
	Password  string        `db:"password"`
	Tokens    *Tokens       `db:"tokens"`
	CreatedAt time.Time     `db:"created_at"`
	DeletedAt *sql.NullTime `db:"deleted_at"`
}

func New(dto request.CreateUsers) Users {
	return Users{
		ID:        uuid.New().String(),
		Email:     dto.Email,
		Name:      dto.Name,
		Password:  dto.Password,
		CreatedAt: time.Now(),
	}
}

func Login(dto request.Login) Users {
	return Users{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (u *Users) HashPassword() error {
	return nil
}

func (u *Users) ComparePassword() error {
	return nil
}
