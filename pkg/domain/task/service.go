package task

import "github.com/GirigiriG/Clean-Architecture-golang/tools"

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
func (s *Service) Create(t *Task) (*Task, error) {

	taskToCreate, err := NewTask(t)
	if err != nil {
		return nil, err
	}

	taskToCreate.ID = tools.CreateUUID()

	err = s.repo.Create(taskToCreate)
	if err != nil {
		return nil, err
	}

	return taskToCreate, nil
}

//DeleteByID delete task record by ID
func (s *Service) DeleteByID(ID string) error {
	if err := s.repo.DeleteByID(ID); err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(t *Task) (*Task, error) {
	toUpdate, err := UpdateTask(t)
	if err != nil {
		return nil, err
	}
	if err = s.repo.Update(t); err != nil {
		return nil, err
	}
	return toUpdate, nil
}

//FindTaskByID find task record by ID
func (s *Service) FindByID(ID string) (*Task, error) {
	result, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//FindAllTaskByProjectID return all task by project the task IDs
func (s *Service) FindAllByProjectID(IDs []string) ([]Task, error) {
	results, err := s.repo.FindAllByProjectID(IDs)
	if err != nil {
		return nil, err
	}
	return results, nil
}
