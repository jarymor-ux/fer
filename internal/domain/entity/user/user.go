package user

import "github.com/jarymor-ux/fer/internal/domain/types"

type User struct {
	ID    types.UserID
	Phone types.PhoneNumber
}
