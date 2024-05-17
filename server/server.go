package server

import (
	"github.com/labstack/echo/v4"
	"github.com/malikilamalik/freemont/config"
)

type Server struct {
	app *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	return &Server{
		app: e,
	}
}

func (s *Server) Run() error {
	return s.app.Start(config.ServerPort())
}
