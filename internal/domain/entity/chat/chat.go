package chat

import (

	"github.com/jarymor-ux/fer/internal/domain/types"
)

type Chat struct {
	chatID               types.ChatID
	firstMessageSenderID types.UserID
	chatType             types.ChatType
	members              types.HashSet[types.UserID]
	history              []types.MessageID
	createdAt            int64
	updatedAt            int64
}

func (c *Chat) Type() types.ChatType {
	return c.chatType
}

func (c *Chat) Members() types.HashSet[types.UserID] {
	return c.members
}

func (c *Chat) History() []types.MessageID {
	return c.history
}

func (c *Chat) CreatedAt() int64 {
	return c.createdAt
}

func (c *Chat) UpdatedAt() int64 {
	return c.updatedAt
}