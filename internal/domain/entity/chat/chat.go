package chat

import (
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/domain/entity/message"
	"github.com/jarymor-ux/fer/internal/domain/entity/types"
)

type Chat struct {
	ChatID, FirstMessageSenderID, Member1ID, Member2ID uuid.UUID
	History                                            types.HashSet[message.Message]
	Conn                                               ws.Conn
}
