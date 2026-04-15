package chat

import (
	"time"

	"github.com/jarymor-ux/fer/internal/domain/types"
)

type Chat struct {
	chatID               types.ChatID
	firstMessageSenderID types.UserID
	chatType             types.ChatType
	members              map[types.UserID]struct{}
	history              []types.MessageID
	createdAt            time.Time
	updatedAt            time.Time
}

func (c *Chat) Type() types.ChatType {
	return c.chatType
}

func (c *Chat) Members() map[types.UserID]struct{} {
	return c.members
}

func (c *Chat) History() []types.MessageID {
	return c.history
}

func (c *Chat) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Chat) UpdatedAt() time.Time{
	return c.updatedAt
}