package message

import (
	"time"

	"braces.dev/errtrace"
	"github.com/google/uuid"
	"github.com/jarymor-ux/fer/internal/domain/errs"
	"github.com/jarymor-ux/fer/internal/domain/types"
	"github.com/jarymor-ux/fer/internal/utils"
)

type Message struct {
	id          types.MessageID
	chatID      types.ChatID
	senderID    types.UserID
	messageText string
	createdAt   time.Time
	updatedAt 	time.Time
	isRead 		bool
}

func (m *Message) MarkRead(){
	if !m.isRead{
		m.isRead = true
	}
}

func(m *Message) EditMessage(newText string) error{
	if newText == "" {
		return errtrace.Wrap(errs.NewError(errs.EmptyMessageError))
	}

	if newText == m.messageText{
		return nil
	}
	m.messageText = newText
	m.updatedAt = time.Now()
	return nil
}

func NewMessage(chatID types.ChatID, senderID types.UserID, text string) (*Message, error) {
	if text == "" {
		return nil, errtrace.Wrap(errs.NewError(errs.EmptyMessageError))
	}

		if chatID == types.ChatID(uuid.Nil) {
		return nil, errtrace.Wrap(errs.NewError(errs.EmptyChatIdError))
	}

		if senderID == types.UserID(uuid.Nil) {
		return nil, errtrace.Wrap(errs.NewError(errs.EmptySenderError))
	}


	return &Message{
		id:          types.MessageID(utils.GenerateUUID()),
		chatID:      chatID,
		senderID:    senderID,
		messageText: text,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		isRead: 	 false,
	}, nil
}

func (m *Message) ID() types.MessageID {
	return m.id
}

func (m *Message) ChatID() types.ChatID {
	return m.chatID
}

func (m *Message) SenderID() types.UserID {
	return m.senderID
}

func (m *Message) Text() string {
	return m.messageText
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Message) UpdatedAt() time.Time {
	return m.updatedAt
}

func (m *Message) IsRead() bool {
	return m.isRead
}