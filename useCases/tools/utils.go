package tools

import (
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

//CreateUUID creates new uuid
func CreateUUID() uuid.UUID {
	return uuid.NewV4()
}

//GetParam gets http params
func GetParam(parmKey string, r *http.Request) string {
	return mux.Vars(r)[parmKey]
}

type ErrorStatus struct {
	Code        int
	Description string
}
