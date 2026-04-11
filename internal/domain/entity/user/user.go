package user

import (
	"github.com/google/uuid"
    ws "github.com/gorilla/websocket"
)

type User struct {
	uid uuid.UUID
	conn *ws.Conn
	Send chan []byte
	Chats any //TODO:implement this
	GroupChats any //TODO:implement this
}
