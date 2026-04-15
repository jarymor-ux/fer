package app

import (
	"github.com/jarymor-ux/fer/internal/io/ws"
	"github.com/labstack/echo/v5"
)

func setAPIRoutes(e *echo.Echo, handlers ws.Handlers) {
	e.GET("/ws", handlers.UpConn) //Убрать в файл api_routes.go func setApiroutes
}
