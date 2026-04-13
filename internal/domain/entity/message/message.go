package message

import (
	"time"

	"github.com/jarymor-ux/fer/internal/domain/types"
	"github.com/jarymor-ux/fer/internal/utils"
)

type Message struct {
	ID          types.MessageID
	ChatID      types.ChatID
	SenderID    types.UserID
	MessageText []byte
	CreatedAt   time.Time
}

func NewMessage(chatID types.ChatID, senderID types.UserID, msg []byte) Message {
	return Message{
		ID:          types.MessageID(utils.GenerateUUID()),
		ChatID:      chatID,
		SenderID:    senderID,
		MessageText: msg,
		CreatedAt:   time.Now(),
	}
}
