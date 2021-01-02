package task

//Repository for task object
type Repository interface {
	Create(*Task) error
	Update(*Task) error
	DeleteByID(string) error
	FindByID(string) (*Task, error)
	FindAllByProjectID([]string) ([]Task, error)
}
