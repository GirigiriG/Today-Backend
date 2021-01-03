package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/project"
	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	projectService *project.Service
	router         *mux.Router
}

func NewProjectHandler(service *project.Service, r *mux.Router) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
		router:         r,
	}
}

func (handler *ProjectHandler) CreateNewProjejct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newProject := &project.Project{}
	err := json.NewDecoder(r.Body).Decode(newProject)

	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	newProject, err = project.NewProject(newProject)
	fmt.Printf("%+v\n", newProject)

	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	err = handler.projectService.CreateNewProjejct(newProject)

	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(newProject)
}

func (handler *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	resutls, err := handler.projectService.GetProjectByID(ID)
	if err != nil {
		w.Write(NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(resutls)

}

func (handler *ProjectHandler) DeleteProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	if err := handler.projectService.DeleteProjectByID(ID); err != nil {
		w.Write(NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
}

func (handler *ProjectHandler) UpdateProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	recordToUpdate := &project.Project{}
	err := json.NewDecoder(r.Body).Decode(recordToUpdate)

	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	resutls, err := handler.projectService.UpdateProjectByID(recordToUpdate)
	if err != nil {
		w.Write(NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(resutls)
}

func (handler *ProjectHandler) HandleRoutes() {
	handler.router.HandleFunc("/project/create", handler.CreateNewProjejct).Methods("GET")
	handler.router.HandleFunc("/project/find/{id}", handler.GetProjectByID).Methods("GET")
	handler.router.HandleFunc("/project/delete/{id}", handler.DeleteProjectByID).Methods("GET")
	handler.router.HandleFunc("/project/update", handler.UpdateProjectByID).Methods("POST")
}
