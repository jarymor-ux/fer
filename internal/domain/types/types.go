package types

import (
	"regexp"
	"strings"

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

func (n PhoneNumber) NormalizePhone() string{
	s := string(n)
	
	s = strings.NewReplacer(" ", "", "-", "", "(", "", ")", "").Replace(s)
	
	if strings.HasPrefix(s, "8") {
		s = "+7" + s[1:]
	} else if !strings.HasPrefix(s, "+7") && len(s) == 10{
		s = "+7" + s
	}
	
	return s
}

func (n PhoneNumber) ValidatePhone() bool {
	normalized := n.NormalizePhone()

	if len(normalized) < 12 {
		return false
	}
	
	re := regexp.MustCompile(`^\+7[3-9]\d{9}$`)
	return re.Match([]byte(normalized))
}


