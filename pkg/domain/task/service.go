package task

import (
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"
)

//Service hold taks and repo
type Service struct {
	repo Repository
	task Task
}

//NewTaskService create a new task service given a repo
func NewTaskService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//Create creates new task record
func (service *Service) Create(t *Task) (*Task, error) {

	taskToCreate, err := NewTask(t)
	if err != nil {
		return nil, err
	}

	taskToCreate.ID = tools.GenerateStringUUID()

	err = service.repo.Create(taskToCreate)
	if err != nil {
		return nil, err
	}

	return taskToCreate, nil
}

//DeleteByID delete task record by ID
func (service *Service) DeleteByID(ID string) error {
	if err := service.repo.DeleteByID(ID); err != nil {
		return err
	}
	return nil
}

//Update : Update task by id
func (service *Service) Update(t *Task) (*Task, error) {
	toUpdate, err := UpdateTask(t)
	if err != nil {
		return nil, err
	}
	if err = service.repo.Update(t); err != nil {
		return nil, err
	}
	return toUpdate, nil
}

//FindByID : Find task record by ID
func (service *Service) FindByID(ID string) (*Task, error) {
	record, err := service.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

//FindAllByProjectID : Get all task by project the task IDs
func (service *Service) FindAllByProjectID(IDs []string) ([]Task, error) {
	records, err := service.repo.FindAllByProjectID(IDs)
	if err != nil {
		return nil, err
	}
	return records, nil
}
