package project

//Repository project repo
type Repository interface {
	GetProjectByID(string) (*Project, error)
	CreateNewProjejct(*Project) error
	DeleteProjectByID(string) error
	UpdateProjectByID(*Project) error
}
