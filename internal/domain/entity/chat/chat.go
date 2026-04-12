package chat

import (
	"github.com/jarymor-ux/fer/internal/domain/types"
)

type Chat struct {
	ChatID               types.ChatID
	FirstMessageSenderID types.UserID
	Type                 types.ChatType
	Members              types.HashSet[types.UserID]
	History              []types.MessageID
	CreatedAt            int64
	UpdatedAt            int64
}
