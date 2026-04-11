package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func (s *Server) Run() error {
	return nil
}
