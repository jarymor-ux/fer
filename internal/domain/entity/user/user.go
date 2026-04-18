package user

import (
	"braces.dev/errtrace"

	"github.com/jarymor-ux/fer/internal/domain/errs"
	"github.com/jarymor-ux/fer/internal/domain/types"
	"github.com/jarymor-ux/fer/internal/utils"
)

type User struct {
	id          types.UserID
	phoneNumber types.PhoneNumber
}

func NewUser(phone types.PhoneNumber) (*User, error) {
	isValid := phone.ValidatePhone()

	if !isValid {
		return nil, errtrace.Wrap(errs.NewError(errs.ValidatePhoneError))
	}

	return &User{
		id:          types.UserID(utils.GenerateUUID()),
		phoneNumber: phone,
	}, nil
}

func (u *User) SetPhone(phone types.PhoneNumber) error {
	isValid := phone.ValidatePhone()

	if !isValid {
		return errtrace.Wrap(errs.NewError(errs.ValidatePhoneError))
	}

	u.phoneNumber = phone

	return nil
}

func (u *User) ID() types.UserID {
	return u.id
}

func (u *User) Phone() types.PhoneNumber {
	return u.phoneNumber
}
