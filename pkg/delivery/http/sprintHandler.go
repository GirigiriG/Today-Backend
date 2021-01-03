package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/sprint"
	"github.com/gorilla/mux"
)

//SprintHandler : holds router (mux.router) and service to interact with db
type SprintHandler struct {
	router  *mux.Router
	service *sprint.Service
}

//NewSprintHandler : handles http routing for sprint record type
func NewSprintHandler(service *sprint.Service, router *mux.Router) *SprintHandler {
	return &SprintHandler{
		service: service,
		router:  router,
	}
}

//Create : create new sprint http route
func (handler *SprintHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := &sprint.Sprint{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		fmt.Println(err.Error())
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	s, err = handler.service.Create(s)

	if err != nil {
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(s)

}

//FindByID : Find sprint by id http route
func (handler *SprintHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)
	if len(ID) != LengthOfUUID {
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	record, err := handler.service.FindByID(ID)
	if err != nil {

		w.Write(NewHTTPError(http.StatusNotFound, "Record not found"))
		return
	}
	json.NewEncoder(w).Encode(record)
}

//UpdateByID : Update sprint by id http route
func (handler *SprintHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := &sprint.Sprint{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	if len(s.ID) != LengthOfUUID {
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	s, err = handler.service.Update(s)
	if err != nil {
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	json.NewEncoder(w).Encode(s)
}

//DeleteByID : delete sprint by id http route
func (handler *SprintHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	if len(ID) != LengthOfUUID {
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	err := handler.service.DeleteByID(ID)
	if err != nil {

		w.Write(NewHTTPError(http.StatusNotFound, "Record not found"))
		return
	}

	w.Write(NewHTTPError(http.StatusOK, "Record "+ID+" successfully deleted"))
}

//HandleRoutes : routing service for sprint record
func (handler *SprintHandler) HandleRoutes() {
	handler.router.HandleFunc("/sprint/create", handler.Create).Methods("GET")
	handler.router.HandleFunc("/sprint/{id}", handler.FindByID).Methods("GET")
	handler.router.HandleFunc("/sprint/update", handler.UpdateByID).Methods("POST")
	handler.router.HandleFunc("/sprint/delete/{id}", handler.DeleteByID).Methods("GET")
}
