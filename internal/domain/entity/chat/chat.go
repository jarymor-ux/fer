package chat

import (
	"strings"
	"sync"
	"time"

	"github.com/jarymor-ux/fer/internal/domain/entity/message"
	"github.com/jarymor-ux/fer/internal/domain/types"
	"github.com/jarymor-ux/fer/internal/utils"
)

type Chat struct {
	chatID               types.ChatID
	firstMessageSenderID types.UserID
	chatType             types.ChatType
	members              map[types.UserID]struct{}
	mu 					 sync.RWMutex
	history              []message.Message
	createdAt            time.Time
	updatedAt            time.Time
}

func NewPrivateChat (creator types.UserID) *Chat {
    return &Chat{
        chatID:               types.ChatID(utils.GenerateUUID()),
        chatType:             types.PrivateChat,
        firstMessageSenderID: creator,
        members:              map[types.UserID]struct{}{creator: {}},
		history: 			  []message.Message{},
        createdAt:            time.Now(),
        updatedAt:            time.Now(),
    }
}

func (c *Chat) AddMessage(msg message.Message){
	c.history = append(c.history, msg)
}

func (c *Chat) hasMember(id types.UserID) bool{
	_,found := c.members[id]
	return found
}

func (c *Chat) HasMember(id types.UserID) bool{
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.hasMember(id)
}

func (c *Chat) AddMember(id types.UserID)bool{
	c.mu.Lock()

	defer c.mu.Unlock()

	if !c.hasMember(id){
	c.members[id] = struct{}{}
	return true
	}
	
	return false
}

func (c *Chat) RemoveMember(id types.UserID)bool{
	c.mu.Lock()

	defer c.mu.Unlock()

	if c.hasMember(id){
	delete(c.members, id)
	return true
	}
	return false
}

func (c *Chat) Type() types.ChatType {
	return c.chatType
}

func (c *Chat) Members() map[types.UserID]struct{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	membersCopy := make(map[types.UserID]struct{}, len(c.members))
	for k, v := range c.members {
		membersCopy[k] = v
	}

	return membersCopy
}

func (c *Chat)GetSubstringMsg(query string) []types.MessageID{
	var res []types.MessageID 
	for _,msg := range c.history{
		if strings.Contains(msg.Text(),query){
			res = append(res, msg.ID())
		}
	}

	return res
}

func (c *Chat)GetMessageByID(id types.MessageID)message.Message{
	for _,msg := range c.history{
		if msg.ID() == id{
			return msg
		}
	}

	return message.Message{}
}

func (c *Chat) GetHistory() []message.Message{
	return c.history
}

func (c *Chat) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Chat) UpdatedAt() time.Time{
	return c.updatedAt
}