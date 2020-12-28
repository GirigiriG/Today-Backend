package task

//Repository for task object
type Repository interface {
	CreateTask(*Task) error
	DeleteTaskByID(string) error
	FindTaskByID(string) (*Task, error)
	FindAllTaskByProjectID([]string) (*[]Task, error)
}
