package project

type Service struct {
	repo Repository
}

func NewProjectService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (service *Service) GetProjectByID(ID string) (*Project, error) {
	result, err := service.repo.GetProjectByID(ID)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (service *Service) CreateNewProjejct(newProject *Project) error {
	err := service.repo.CreateNewProjejct(newProject)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) DeleteProjectByID(ID string) error {
	err := service.repo.DeleteProjectByID(ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateProjectByID(recordToUpdate *Project) (*Project, error) {
	result, err := service.repo.UpdateProjectByID(recordToUpdate)
	if err != nil {
		return nil, err
	}
	return result, nil
}
