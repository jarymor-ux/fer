package http

import (
	"net/http"

	"braces.dev/errtrace"
	ws "github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/io/http/response"
)

type Handlers struct {
	upgraider *ws.Upgrader
}

func newUpgraider() *ws.Upgrader {
	return &ws.Upgrader{
		HandshakeTimeout: 5,
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
		upgraider: newUpgraider(),
	}
}


func (h *Handlers) setUpAndUpgade(w http.ResponseWriter, r *http.Request) error {
	conn, err := h.upgraider.Upgrade(w,r, nil)
	defer conn.Close()
	if err != nil {
		resp, err := response.NewHTTPError("Failed to create connection").GetJSONBytes()
		if err != nil {
			return errtrace.Wrap(err)
		}
		w.Write(resp)
		return errtrace.Wrap(err)
	}
	//TODO:Нужно собирать данные пользака по типу айпи и тд
	return nil
}
