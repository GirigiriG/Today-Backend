package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/delivery"

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

func (handler *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {

}

func (handler *ProjectHandler) CreateNewProjejct(w http.ResponseWriter, r *http.Request) {
	newProject := &project.Project{}
	err := json.NewDecoder(r.Body).Decode(newProject)

	if err != nil {
		panic(err)
	}

	newProject, err = project.NewProject(newProject)

	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	err = handler.projectService.CreateNewProjejct(newProject)

	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}

	newProject, err = handler.projectService.GetProjectByID(newProject.ID)
	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusBadRequest, err.Error()))
	}

	json.NewEncoder(w).Encode(newProject)

}

func (handler *ProjectHandler) DeleteProjectByID(w http.ResponseWriter, r *http.Request) {

}

func (handler *ProjectHandler) UpdateProjectByID(w http.ResponseWriter, r *http.Request) {

}

func (handler *ProjectHandler) HandleProjectRoutes() {
	handler.router.HandleFunc("/project/create", handler.CreateNewProjejct).Methods("GET")
}
