package routes

import (
	"net/http"

	middleware "github.com/GirigiriG/Clean-Architecture-golang/middlerware"
	googleoauth "github.com/GirigiriG/Clean-Architecture-golang/pkg"
	"github.com/gorilla/mux"
)

type unprotectedRoutes struct {
	router *mux.Router
}

func NewUnprotectedRoutesHandler(r *mux.Router) *unprotectedRoutes {
	return &unprotectedRoutes{
		router: r,
	}
}

func (handler *unprotectedRoutes) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	url := googleoauth.GoogleOauthConfig.AuthCodeURL(googleoauth.RandomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	w.Write([]byte(middleware.GenerateJWToken()))
}

func (handler *unprotectedRoutes) InitProtectedRoutes() {
	handler.router.HandleFunc("/login", handler.login).Methods("GET")
}
