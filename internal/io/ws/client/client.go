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

func NewClient(ip net.IP, clientid types.ClientID, userid types.UserID, conn *ws.Conn) *Client {
	return &Client{
		ClientIP: ip,
		ClientID: clientid,
		UserID:   userid,
		Conn:     conn,
		Send:     make(chan []byte),
	}
}
