package types

import (
	"regexp"
	"strings"

	"braces.dev/errtrace"
	"github.com/google/uuid"
	"github.com/hashicorp/go-set/v3"
	"github.com/jarymor-ux/fer/internal/domain/errs"
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

func (n PhoneNumber) NormalizePhone() (string, error) {
	s := string(n)

	s = strings.NewReplacer(" ", "", "-", "", "(", "", ")", "").Replace(s)

	switch {
	case strings.HasPrefix(s, "8") && len(s) == 11:
		return "+7" + s[1:], nil

	case !strings.HasPrefix(s, "+7") && len(s) == 10:
		return "+7" + s, nil

	case strings.HasPrefix(s, "+7") && len(s) == 12:
		return s, nil

	default:
		return "", errtrace.Wrap(errs.NewError(errs.ValidatePhoneError))
	}
}

func (n PhoneNumber) ValidatePhone() (bool, error) {
	normalized, err := n.NormalizePhone()
	if err != nil {
		return false, errtrace.Wrap(err)
	}

	re := regexp.MustCompile(`^\+7[3-9]\d{9}$`)

	res := re.Match([]byte(normalized))
	if res {
		return res, nil
	}

	return res, errtrace.Wrap(errs.NewError(errs.ValidatePhoneError))
}
