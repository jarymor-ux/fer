package http

import (
	"github.com/labstack/echo/v5"
)

func (h *Handlers) UpToWs(c *echo.Context) error {
	_, err := h.setUpAndUpgade(c.Response(), c.Request())
	return err //TODO
}
