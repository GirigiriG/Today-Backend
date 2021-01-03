package project

//Repository project repo
type Repository interface {
	FindByID(string) (*Project, error)
	Create(*Project) error
	DeleteByID(string) error
	UpdateByID(*Project) error
}
