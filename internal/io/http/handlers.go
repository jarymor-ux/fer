package http

import (
	"net/http"
	"time"

	"braces.dev/errtrace"
	ws "github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/io/http/response"
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

func (h *Handlers) setUpAndUpgade(w http.ResponseWriter, r *http.Request) (*ws.Conn, error) {
	conn, uerr := h.upgrader.Upgrade(w, r, nil)
	if uerr != nil {
		resp, err := response.NewHTTPError("Failed to create connection", http.StatusInternalServerError).
			GetJSONBytes()
		if err != nil {
			return nil, errtrace.Wrap(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)

		if _, err = w.Write(resp); err != nil {
			return nil, errtrace.Wrap(err)
		}

		return nil, errtrace.Wrap(uerr)
	}
	//TODO:Нужно собирать данные пользака по типу айпи и тд
	return conn, nil
}
