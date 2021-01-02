package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/delivery"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"
	"github.com/gorilla/mux"
)

//TaskHandler struct
type TaskHandler struct {
	taskService *task.Service
	router      *mux.Router
}

//NewTaskHandler requires http router and service
func NewTaskHandler(service *task.Service, router *mux.Router) *TaskHandler {
	return &TaskHandler{
		taskService: service,
		router:      router,
	}
}

func (handler *TaskHandler) createNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	t := &task.Task{}
	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusBadRequest, "Bad request."))
		return
	}

	t, err = handler.taskService.Create(t)
	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusBadRequest, "Bad request."))
		return
	}

	json.NewEncoder(w).Encode(t)
}

func (handler *TaskHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := &task.Task{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusBadRequest, "Bad request."))
		return
	}

	t, err = handler.taskService.FindByID(t.ID)

	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		w.Write(delivery.NewHttpError(http.StatusNotFound, "Record not found."))
		return
	}

	json.NewEncoder(w).Encode(t)

}

func (handler *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (handler *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	if err := handler.taskService.DeleteByID(ID); err != nil {
		w.Write(delivery.NewHttpError(http.StatusNotFound, "Record not found."))
		return
	}
	w.Write(delivery.NewHttpError(http.StatusOK, "Ok"))
}

//FindAllTaskByProjectID get all task from slice of project Ids
func (handler *TaskHandler) FindAllTaskByProjectID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type projectIds struct {
		Ids []string
	}
	defer r.Body.Close()

	taskProjectIds := &projectIds{}
	if err := json.NewDecoder(r.Body).Decode(taskProjectIds); err != nil {
		w.Write(delivery.NewHttpError(http.StatusInternalServerError, err.Error()))
	}

	tasks, err := handler.taskService.FindAllByProjectID(taskProjectIds.Ids)
	if err != nil {
		w.Write(delivery.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

//HandleTaskRoutes all routing for task struct
func (handler *TaskHandler) HandleTaskRoutes() {
	handler.router.HandleFunc("/task/create", handler.createNewTask).Methods("GET")
	handler.router.HandleFunc("/task/{id}", handler.FindByID).Methods("GET")
	handler.router.HandleFunc("/task/creaupdatete", handler.Update).Methods("GET")
	handler.router.HandleFunc("/task/delete/{id}", handler.Delete).Methods("GET")
	handler.router.HandleFunc("/task/project/{id}", handler.FindAllTaskByProjectID).Methods("GET")
}
