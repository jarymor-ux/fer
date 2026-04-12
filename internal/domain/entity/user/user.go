package user

import (
	"github.com/jarymor-ux/fer/internal/domain/entity/types"
)

type User struct {
	ID  types.UserID
	Phone types.PhoneNumber
	Chats []types.ChatID
}
