package delivery

import (
	"encoding/json"
)

type httpError struct {
	Code        int
	Description string
}

func NewHttpError(code int, desc string) []byte {
	httpErr := httpError{Code: code, Description: desc}
	bs, err := json.Marshal(httpErr)
	if err != nil {
		panic(err)
	}
	return bs
}
