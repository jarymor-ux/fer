package message

import (
	"time"

	"github.com/jarymor-ux/fer/internal/domain/types"
)

type Message struct {
	ID          types.MessageID
	ChatID      types.ChatID
	SenderID    types.UserID
	MessageText string
	CreatedAt   time.Time
}
