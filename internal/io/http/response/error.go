package response

import (
	"encoding/json"

	"braces.dev/errtrace"
)

type HTTPError struct {
	Msg string `json:"error_msg"`
}

func NewHTTPError(msg string) *HTTPError {
	return &HTTPError{
		Msg: msg,
	}
}

func (h HTTPError) GetJSONBytes() ([]byte, error) {
	bytes, err := json.Marshal(h)
	if err != nil {
		return []byte{}, errtrace.Wrap(err)
	}
	return bytes, err
}
