package types

import (
	"regexp"

	"github.com/google/uuid"
	"github.com/hashicorp/go-set/v3"
)

type ChatType byte

const (
	PrivateChat ChatType = iota
	GroupChat
)

type (
	PhoneNumber    []byte
	UserID         uuid.UUID
	ClientID       uuid.UUID
	MessageID      uuid.UUID
	ChatID         uuid.UUID
	HashSet[t any] set.HashSet[t, string]
)

func (n PhoneNumber) ValidatePhone() bool {
	if len(n) < 11 {
		return false
	}
	re := regexp.MustCompile(`^(?:\+7|8)9\d{9}$`)
	return re.Match(n)
}

