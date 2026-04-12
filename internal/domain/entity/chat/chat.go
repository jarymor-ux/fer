package chat

import (
	"github.com/google/uuid"
	"github.com/jarymor-ux/fer/internal/domain/types"
)

type Chat struct {
	ChatID, FirstMessageSenderID uuid.UUID
	Type                         types.ChatType
	Members                      types.HashSet[types.UserID]
	History                      []types.MessageID
	CreatedAt                    int64
	UpdatedAt                    int64
}
