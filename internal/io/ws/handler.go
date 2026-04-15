package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/io/ws/upgrader"
	"github.com/labstack/echo/v5"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return new(Handlers)
}

func (h *Handlers) UpConn(e *echo.Context) error {
	conn, err := upgrader.New().Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		e.JSON(http.StatusInternalServerError, struct {
			Msg string `json:"msg"`
		}{
			Msg: "Внутренняя ошибка сервера. Попробуйте позже.",
		})
	}

	h.proceed(conn, e.RealIP())

	return nil
}

func (h *Handlers) proceed(conn *websocket.Conn, ip string) error {
	req := new(LoginRequest)
	for {
		conn.ReadJSON(req)
		conn.WriteJSON(UserResponse{
			Username: req.Username,
			IP:       ip,
		})
	}
}

type LoginRequest struct {
	Username string `json:"username"`
}

type UserResponse struct {
	Username string `json:"username"`
	IP       string `json:"IP"`
}
