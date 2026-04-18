package errs

import "errors"

type ErrorText string

const (
	ValidatePhoneError             ErrorText = "Phone number validation error"
	FailedToUpgradeHTTPToWebSocket ErrorText = "Failed to upgrage connection"
	EmptyMessageError 			   ErrorText = "Empty message error"
	EmptySenderError			   ErrorText = "Emtpy sender error"
	EmptyChatIdError			   ErrorText = "Emtpy chatid error"
	MessageNotFoundError		   ErrorText = "No messages found in history"
)

func NewError(s ErrorText) error {
	return errors.New(string(s))
}
