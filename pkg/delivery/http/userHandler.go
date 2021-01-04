package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

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

//Create : create a new user record
func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	record := &user.User{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad Request"))
		return
	}

	u, err := handler.userService.Create(record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(u)

}

//Update register user service and http router
func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	record := &user.User{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad Request"))
		return
	}

	if len(record.ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	u, err := handler.userService.UpdateByID(record)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(u)
}

//FindByID : get user by id
func (handler *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	// if len(ID) != LengthOfUUID {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
	// 	return
	// }

	u, err := handler.userService.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (handler *UserHandler) deleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	if len(ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	err := handler.userService.DeleteByID(ID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
		return
	}
}

//HandleRoutes handler for user
func (handler *UserHandler) HandleRoutes() {
	handler.router.HandleFunc("/user/create", handler.Create).Methods("GET")
	handler.router.HandleFunc("/user/find/{id}", handler.FindByID).Methods("GET")
	handler.router.HandleFunc("/user/update", handler.Update).Methods("GET")
	handler.router.HandleFunc("/user/delete/{id}", handler.deleteByID).Methods("GET")
}
