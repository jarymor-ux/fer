package response

import (
	"encoding/json"

	"braces.dev/errtrace"
)

type HTTPError struct {
	Msg    string `json:"error_msg"`
	Status int    `json:"status"`
}

func NewHTTPError(msg string, status int) *HTTPError {
	return &HTTPError{
		Msg:    msg,
		Status: status,
	}
}

func (h HTTPError) GetJSONBytes() ([]byte, error) {
	bytes, err := json.Marshal(h)
	if err != nil {
		return []byte{}, errtrace.Wrap(err)
	}

	return bytes, err
}
