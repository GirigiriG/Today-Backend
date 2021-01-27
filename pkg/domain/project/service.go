package project

import "github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

//Service : holds repository
type Service struct {
	repo Repository
}
//NewProjectService : requires concrete implmentation of Repository interface
func NewProjectService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//FindByID : find project by id
func (service *Service) FindByID(ID string) (*Project, error) {
	record, err := service.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return record, nil

}

//Create : create new project 
func (service *Service) Create(newProject *Project) error {
	newProject.ID = tools.GenerateStringUUID()
	err := service.repo.Create(newProject)
	if err != nil {
		return err
	}

	return nil
}

//DeleteByID : delete project by id
func (service *Service) DeleteByID(ID string) error {
	err := service.repo.DeleteByID(ID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateByID : update project by id
func (service *Service) UpdateByID(recordToUpdate *Project) (*Project, error) {
	recordToUpdate, err := UpdateProject(recordToUpdate)
	if err != nil {
		return nil, err
	}

	err = service.repo.UpdateByID(recordToUpdate)
	if err != nil {
		return nil, err
	}
	return recordToUpdate, nil
}
