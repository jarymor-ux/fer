package client

import (
	"net"

	ws "github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/domain/types"
)

type Client struct {
	ClientIP net.IP
	ClientID types.ClientID
	UserID   types.UserID
	Conn     *ws.Conn
	Send     chan []byte
}
