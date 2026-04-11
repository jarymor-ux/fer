package message

import (
	"github.com/google/uuid"
)

type Message struct {
	MessageID, SenderID, ReciverID uuid.UUID
	MessageText []byte
	Time int64
}