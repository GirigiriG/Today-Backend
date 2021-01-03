package delivery

import (
	"encoding/json"
)

type httpError struct {
	Code        int
	Description string
}

//NewHTTPError : returns http status and error message
func NewHTTPError(code int, desc string) []byte {
	httpErr := httpError{Code: code, Description: desc}
	bs, err := json.Marshal(httpErr)
	if err != nil {
		panic(err)
	}
	return bs
}

const (
	//LengthOfUUID : The leinght of a UUID
	LengthOfUUID = 36
)
