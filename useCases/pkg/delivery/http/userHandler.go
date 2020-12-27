package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	middleware "github.com/GirigiriG/Clean-Architecture-golang/middlerware"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/user"
	"github.com/gorilla/mux"
)

//UserHandler holds user service and router
type UserHandler struct {
	userService *user.Service
	router      *mux.Router
}

//NewUserHandler register user service and http router
func NewUserHandler(userCase *user.Service, r *mux.Router) *UserHandler {

	return &UserHandler{
		userService: userCase,
		router:      r,
	}
}

func (handler *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)
	fmt.Println(ID)
	u, err := handler.userService.GetUserByID(ID)

	if err != nil {
		handleError(err, w)
		return
	}
	user, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	w.Write(user)
}

func (handler *UserHandler) deleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	err := handler.userService.DeleteUserByID(ID)

	if err != nil {
		handleError(err, w)
		return
	}
}

func (handler *UserHandler) createNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userOutCh := make(chan []byte, 1)

	var u user.User

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &u)

	go func(u *user.User) {
		u, err := handler.userService.CreateNewUser(u)
		if err != nil {
			userOutCh <- []byte(err.Error())
		}

		resp, err := json.Marshal(u)

		if err != nil {
			panic(err)
		}

		userOutCh <- resp

	}(&u)

	w.Write(<-userOutCh)

}

func handleError(e error, w http.ResponseWriter) {
	errorMsg, _ := json.Marshal(e.Error())
	w.Write(errorMsg)

}

//HandleUserRoutes handler for user
func (handler *UserHandler) HandleUserRoutes() {
	handler.router.HandleFunc("/user/create", middleware.IsAuthorized(handler.createNewUser)).Methods("GET")
	handler.router.HandleFunc("/user/{id}", middleware.IsAuthorized(handler.getUserByID)).Methods("GET")
	handler.router.HandleFunc("/user/delete/{id}", handler.deleteUserByID).Methods("GET")
}
