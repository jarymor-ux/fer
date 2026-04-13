package http

import (
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"
)

type Handlers struct {
	upgrader *ws.Upgrader
}

func newUpgrader() *ws.Upgrader {
	return &ws.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   2048, // TODO:пересмотреть размеры буферов
		WriteBufferSize:  2048,
		//WriteBufferPool:  nil, //Написать свою реализацию или же использовать готовый
		//Subprotocols:     []string{}, //TODO:разобпаться что это и зачем
		//Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		//	panic("TODO")
		//},
		CheckOrigin: func(r *http.Request) bool {
			return true // TODO:переписать в будущем
		},
		EnableCompression: true,
	}
}

func NewHandlers() *Handlers {
	return &Handlers{
		upgrader: newUpgrader(),
	}
}
