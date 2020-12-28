package task

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

//CreateTask creates new task record
func (s *Service) CreateTask(task *Task) error {
	newTask, err := NewTask(task)
	if err != nil {
		return err
	}
	err = s.repo.CreateTask(newTask)

	if err != nil {
		return err
	}

	return nil
}

//DeleteTaskByID delete task record by ID
func (s *Service) DeleteTaskByID(ID string) error {
	if err := s.repo.DeleteTaskByID(ID); err != nil {
		return err
	}
	return nil
}

//FindTaskByID find task record by ID
func (s *Service) FindTaskByID(ID string) (*Task, error) {
	result, err := s.repo.FindTaskByID(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//FindAllTaskByProjectID return all task by project the task IDs
func (s *Service) FindAllTaskByProjectID(IDs []string) (*[]Task, error) {
	results, err := s.repo.FindAllTaskByProjectID(IDs)
	if err != nil {
		return nil, err
	}

	return results, nil
}
