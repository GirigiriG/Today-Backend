package routes

import (
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/auth"
	"github.com/gorilla/mux"
)

//UnprotectedRoutes : holds a router
type UnprotectedRoutes struct {
	router *mux.Router
}

//NewUnprotectedRoutesHandler : struct for unprotected route
func NewUnprotectedRoutesHandler(r *mux.Router) *UnprotectedRoutes {
	return &UnprotectedRoutes{
		router: r,
	}
}

func (handler *UnprotectedRoutes) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//url := googleoauth.GoogleOauthConfig.AuthCodeURL(googleoauth.RandomState)
	//http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	w.Write([]byte(auth.GenerateJWToken()))
}

//InitProtectedRoutes call login handler return jwt
func (handler *UnprotectedRoutes) InitProtectedRoutes() {
	handler.router.HandleFunc("/login", handler.login).Methods("GET")
}
