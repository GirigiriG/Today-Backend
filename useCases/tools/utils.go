package tools

import (
	"net/http"

	"github.com/gorilla/mux"
)

//CreateUUID creates new uuid

//GetParam gets http params
func GetParam(parmKey string, r *http.Request) string {
	return mux.Vars(r)[parmKey]
}

type ErrorStatus struct {
	Code        int
	Description string
}
