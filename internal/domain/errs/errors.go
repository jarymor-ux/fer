package errs

import "errors"

type ErrorText string

const (
	ValidatePhoneError ErrorText = "Phone number validation error"
)

func NewError(s ErrorText) error {
	return errors.New(string(s))
}
