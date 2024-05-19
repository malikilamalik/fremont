package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/malikilamalik/freemont/config"
	v1 "github.com/malikilamalik/freemont/internal/delivery/http/v1"
	"github.com/malikilamalik/freemont/pkg/database/postgresql"
)

type Server struct {
	app       *echo.Echo
	db        *sqlx.DB
	validator *validator.Validate
}

func NewServer() *Server {
	e := echo.New()
	db := postgresql.New()
	validate := validator.New()

	return &Server{
		app:       e,
		db:        db,
		validator: validate,
	}
}

func (s *Server) Run() error {

	//Init Router for v1
	v1.Init(s.app, s.db, s.validator)

	return s.app.Start(config.ServerPort())
}
