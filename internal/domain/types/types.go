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

func (n PhoneNumber) NormalizePhone() (error, string) {
	s := string(n)
	
	s = strings.NewReplacer(" ", "", "-", "", "(", "", ")", "").Replace(s)
	
	if strings.HasPrefix(s, "8") && len(s) == 11 {
		s = "+7" + s[1:]
		return nil, s
	} else if !strings.HasPrefix(s, "+7") && len(s) == 10{
		s = "+7" + s
		return nil, s
	} else if strings.HasPrefix(s, "+7") && len(s) == 12 {
		return nil, s
	}
	
	return errtrace.Wrap(errs.NewError(errs.ValidatePhoneError)), ""
}

func (n PhoneNumber) ValidatePhone() (error ,bool) {
	err, normalized := n.NormalizePhone()
	if err != nil {
		return errtrace.Wrap(err), false
	}

	re := regexp.MustCompile(`^\+7[3-9]\d{9}$`)

    res := re.Match([]byte(normalized))
	if res {
		return nil, res
	}

	return errtrace.Wrap(errs.NewError(errs.ValidatePhoneError)), res
}

