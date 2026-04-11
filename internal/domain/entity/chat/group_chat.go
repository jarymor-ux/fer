package chat

import (
	ws "github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/domain/entity/message"
	"github.com/jarymor-ux/fer/internal/domain/entity/types"
	"github.com/jarymor-ux/fer/internal/domain/entity/user"
)

type GroupChat struct {
	ChatID  types.ChatID
	Members types.HashSet[user.User]
	History types.HashSet[message.Message]
	Conn    ws.Conn
}
