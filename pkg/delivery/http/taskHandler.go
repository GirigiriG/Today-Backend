package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"
	"github.com/gorilla/mux"
)

//TaskHandler struct
type TaskHandler struct {
	service *task.Service
	router  *mux.Router
}

//NewTaskHandler requires http router and service
func NewTaskHandler(service *task.Service, router *mux.Router) *TaskHandler {
	return &TaskHandler{
		service: service,
		router:  router,
	}
}

func (handler *TaskHandler) createNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	t := &task.Task{}
	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, "Bad request."))
		return
	}

	t, err = handler.service.Create(t)
	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, "Bad request."))
		return
	}

	json.NewEncoder(w).Encode(t)
}

//FindByID : task by record id
func (handler *TaskHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := &task.Task{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, "Bad request."))
		return
	}

	t, err = handler.service.FindByID(t.ID)

	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		w.Write(NewHttpError(http.StatusNotFound, "Record not found."))
		return
	}

	json.NewEncoder(w).Encode(t)
}

//Update : update task record
func (handler *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := &task.Task{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.Write(NewHttpError(http.StatusBadRequest, "Bad request"))
		return
	}

	t, err = handler.service.Update(t)
	if err != nil {
		w.Write(NewHttpError(http.StatusNotFound, "Record not found."))
		return
	}

	json.NewEncoder(w).Encode(t)
}

//Delete :  task record by id
func (handler *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)

	if err := handler.service.DeleteByID(ID); err != nil {
		w.Write(NewHttpError(http.StatusNotFound, "Record not found."))
		return
	}
	w.Write(NewHttpError(http.StatusOK, "Ok"))
}

//FindAllTaskByProjectID get all task from slice of project Ids
func (handler *TaskHandler) FindAllTaskByProjectID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := tools.GetParam("id", r)
	type projectIds struct {
		Ids []string
	}

	taskProjectIds := &projectIds{}
	defer r.Body.Close()

	if len(ID) != 0 {
		tasks, err := handler.service.FindAllByProjectID([]string{ID})
		if err != nil {
			w.Write(NewHttpError(http.StatusInternalServerError, err.Error()))
			return
		}
		json.NewEncoder(w).Encode(tasks)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(taskProjectIds); err != nil {
		w.Write(NewHttpError(http.StatusInternalServerError, err.Error()))
	}

	tasks, err := handler.service.FindAllByProjectID(taskProjectIds.Ids)
	if err != nil {
		w.Write(NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

//HandleRoutes : all routing for task struct
func (handler *TaskHandler) HandleRoutes() {
	handler.router.HandleFunc("/task/create", handler.createNewTask).Methods("GET")
	handler.router.HandleFunc("/task/find/{id}", handler.FindByID).Methods("GET")
	handler.router.HandleFunc("/task/update", handler.Update).Methods("POST")
	handler.router.HandleFunc("/task/delete/{id}", handler.Delete).Methods("GET")
	handler.router.HandleFunc("/task/projects/{id}", handler.FindAllTaskByProjectID).Methods("GET")
	handler.router.HandleFunc("/task/projects/", handler.FindAllTaskByProjectID).Methods("GET")
}
