package handlers

import (
	ws "github.com/gorilla/websocket"
)

type Handlers struct {
	upgraider *ws.Upgrader
}


