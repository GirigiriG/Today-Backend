package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/project"
	"github.com/gorilla/mux"
)

//ProjectHandler : Holds project service and router (mux)
type ProjectHandler struct {
	projectService *project.Service
	router         *mux.Router
}

//NewProjectHandler : Hold project service and router (mux)
func NewProjectHandler(service *project.Service, r *mux.Router) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
		router:         r,
	}
}

//Create : Create a new project
func (handler *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newProject := &project.Project{}
	err := json.NewDecoder(r.Body).Decode(newProject)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}

	newProject, err = project.NewProject(newProject)
	fmt.Printf("%+v\n", newProject)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}
	err = handler.projectService.Create(newProject)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newProject)
}

//FindByID : Find project by Id
func (handler *ProjectHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)
	if len(ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	resutls, err := handler.projectService.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resutls)
}

//DeleteByID : delete project by Id
func (handler *ProjectHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)
	if len(ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	if err := handler.projectService.DeleteByID(ID); err != nil {
		w.Write(NewHTTPError(http.StatusInternalServerError, err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

//UpdateByID : Update project by Id
func (handler *ProjectHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	record := &project.Project{}
	err := json.NewDecoder(r.Body).Decode(record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}

	if len(record.ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	resutls, err := handler.projectService.UpdateByID(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(NewHTTPError(http.StatusInternalServerError, err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resutls)
}

//HandleRoutes router handler for project
func (handler *ProjectHandler) HandleRoutes() {
	handler.router.HandleFunc("/project/create", handler.Create).Methods("GET")
	handler.router.HandleFunc("/project/find/{id}", handler.FindByID).Methods("GET")
	handler.router.HandleFunc("/project/delete/{id}", handler.DeleteByID).Methods("GET")
	handler.router.HandleFunc("/project/update", handler.UpdateByID).Methods("POST")
}
