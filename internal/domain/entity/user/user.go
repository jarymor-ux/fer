package user

import (
	"braces.dev/errtrace"
	"github.com/jarymor-ux/fer/internal/domain/errs"
	"github.com/jarymor-ux/fer/internal/domain/types"
	"github.com/jarymor-ux/fer/internal/utils"
)

type User struct {
	id  types.UserID
	phoneNumber types.PhoneNumber
}

func NewUser(phone types.PhoneNumber)(*User,error){
	res,err  := phone.ValidatePhone()

	if err != nil{
		return nil,errtrace.Wrap(err)
	}

	if !res{
		return nil,errtrace.Wrap(errs.NewError(errs.ValidatePhoneError))
	}
	
	return &User{
		id: types.UserID(utils.GenerateUUID()),
		phoneNumber: phone,
	},nil

}

func (u *User)ChangePhone(phone types.PhoneNumber)error {
	res,err  := phone.ValidatePhone()

	
	if err != nil{
		return errtrace.Wrap(err)
	}

	if !res{
		return errtrace.Wrap(errs.NewError(errs.ValidatePhoneError))
	}
	
	u.phoneNumber = phone
	return nil
}

func (u *User) ID() types.UserID{
	return u.id
}

func (u *User) Phone() types.PhoneNumber{
	return u.phoneNumber
}