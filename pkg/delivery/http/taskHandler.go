package delivery

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	defer r.Body.Close()

	t := &task.Task{}
	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request."))
		return
	}

	t, err = handler.service.Create(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(t)
}

//FindByID : task by record id
func (handler *TaskHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID := tools.GetParam("id", r)

	t, err := handler.service.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

//Update : update task record
func (handler *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	t := &task.Task{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	if len(t.ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	t, err = handler.service.Update(t)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

//Delete :  task record by id
func (handler *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID := tools.GetParam("id", r)

	if len(ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}
	if err := handler.service.DeleteByID(ID); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(NewHTTPError(http.StatusOK, "Ok"))
}

//FindAllTaskByProjectID get all task from slice of project Ids
func (handler *TaskHandler) FindAllTaskByProjectID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ID := tools.GetParam("id", r)
	type projectIds struct {
		Ids []string
	}

	if len(ID) != LengthOfUUID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(NewHTTPError(http.StatusBadRequest, "Bad request"))
		return
	}

	defer r.Body.Close()

	if len(ID) != 0 {
		records, err := handler.service.FindAllByProjectID([]string{ID})

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(NewHTTPError(http.StatusNotFound, err.Error()))
			return
		}

		tasks, err := json.Marshal(records)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(tasks)
	}
}

//HandleRoutes : all routing for task struct
func (handler *TaskHandler) HandleRoutes() {
	handler.router.HandleFunc("/task/create", handler.createNewTask).Methods("POST")
	handler.router.HandleFunc("/task/find/{id}", handler.FindByID).Methods("GET")
	handler.router.HandleFunc("/task/update", handler.Update).Methods("POST")
	handler.router.HandleFunc("/task/delete/{id}", handler.Delete).Methods("GET")
	handler.router.HandleFunc("/task/project/{id}", handler.FindAllTaskByProjectID).Methods("GET")
}
