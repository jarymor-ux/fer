package http

import ws "github.com/gorilla/websocket"

type Handlers struct {
	upgraider *ws.Upgrader
}

func NewHandlers(up *ws.Upgrader) *Handlers {
	return &Handlers{
		upgraider: up,
	}
}
