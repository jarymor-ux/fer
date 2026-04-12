package http

import (
	"github.com/labstack/echo/v5"
)

func (h *Handlers) UpToWs(c *echo.Context) error {
	h.setUpAndUpgade(c.Response(),c.Request())
	return nil //TODO
}

