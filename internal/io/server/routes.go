package server

import (
	"github.com/jarymor-ux/fer/internal/io/server/handlers"
	"github.com/labstack/echo/v4"
)

func setRoutes(e *echo.Echo, h handlers.Handlers) error {

	e.GET("/ws", h.Ws)

	return nil
}
