package upgrader

import (
	"github.com/gorilla/websocket"
)

func New() *websocket.Upgrader {
	return &websocket.Upgrader{
		//HandshakeTimeout: 0,
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		//WriteBufferPool:  nil,
		//Subprotocols:     []string{},
		//Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		//	panic("TODO")
		//},
		//CheckOrigin: func(r *http.Request) bool {
		//	panic("TODO")
		//},
		EnableCompression: false,
	}
}
