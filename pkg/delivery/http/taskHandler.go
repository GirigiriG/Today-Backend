package delivery

import (
	"encoding/json"
	"net/http"

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

	var newTask *task.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		panic(err.Error())
	}

	data, err := task.NewTask(newTask)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := handler.taskService.CreateTask(data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

//FindAllTaskByProjectID get all task from slice of project Ids
func (handler *TaskHandler) FindAllTaskByProjectID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type projectIds struct {
		ProjectIDs []string
	}
	defer r.Body.Close()

	taskProjectIds := &projectIds{}
	if err := json.NewDecoder(r.Body).Decode(taskProjectIds); err != nil {
		w.Write([]byte(err.Error()))
	}

	tasks, err := handler.taskService.FindAllTaskByProjectID(taskProjectIds.ProjectIDs)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(tasks)

}

//HandleTaskRoutes all routing for task struct
func (handler *TaskHandler) HandleTaskRoutes() {
	handler.router.HandleFunc("/task/create", handler.createNewTask).Methods("GET")
	handler.router.HandleFunc("/task/project", handler.FindAllTaskByProjectID).Methods("GET")
}
