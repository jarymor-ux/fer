package types

import (
	"regexp"
	"strings"

	
	"github.com/google/uuid"
)

type ChatType byte

const (
	PrivateChat ChatType = iota
	GroupChat
)

type (
	PhoneNumber    string
	UserID         uuid.UUID
	ClientID       uuid.UUID
	MessageID      uuid.UUID
	ChatID         uuid.UUID
)

func (n PhoneNumber) NormalizePhone() string {
	s := string(n)

	s = strings.NewReplacer(" ", "", "-", "", "(", "", ")", "").Replace(s)

	switch {
	case strings.HasPrefix(s, "8") && len(s) == 11:
		return "+7" + s[1:]

	case !strings.HasPrefix(s, "+7") && len(s) == 10:
		return "+7" + s

	case strings.HasPrefix(s, "+7") && len(s) == 12:
		return s

	default:
		return ""
	}
}

func (n PhoneNumber) ValidatePhone() bool {
	normalized:= n.NormalizePhone()

	re := regexp.MustCompile(`^\+7[3-9]\d{9}$`)

	res := re.Match([]byte(normalized))
	if res {
		return res
	}

	return res
}
