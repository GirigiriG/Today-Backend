package tools

import (
	"net/http"

	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

//CreateUUID creates new uuid
func GenerateStringUUID() string {
	id := uuid.New()
	return id.String()
}

//GetParam gets http params
func GetParam(parmKey string, r *http.Request) string {
	return mux.Vars(r)[parmKey]
}

type ErrorStatus struct {
	Code        int
	Description string
}
