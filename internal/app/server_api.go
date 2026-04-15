package app

import (
	"github.com/jarymor-ux/fer/internal/io/ws"
	"github.com/labstack/echo/v5"
)

func newWsServer() *echo.Echo {
	e := echo.New()
	handlers := ws.NewHandlers()
	setAPIRoutes(e, *handlers)

	return e
}
