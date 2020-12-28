package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"
	"github.com/gorilla/mux"
)

type TaskHandler struct {
	taskService *task.Service
	router      *mux.Router
}

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
	}

	if err := handler.taskService.CreateTask(data); err != nil {
		panic(err)
	}
}

func (handler *TaskHandler) HandleTaskRoutes() {
	handler.router.HandleFunc("/task/create", handler.createNewTask).Methods("GET")
}
