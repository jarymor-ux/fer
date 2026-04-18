package errs

import "errors"

type ErrorText string

const (
	ValidatePhoneError             ErrorText = "Phone number validation error"
	FailedToUpgradeHTTPToWebSocket ErrorText = "Failed to upgrage connection"
	EmptyMessageError              ErrorText = "Empty message error"
	EmptySenderError               ErrorText = "Empty sender error"
	EmptyChatIDError               ErrorText = "Empty chatid error"
	MessageNotFoundError           ErrorText = "Message was not found in history"
)

func NewError(s ErrorText) error {
	return errors.New(string(s))
}
