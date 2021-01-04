package sprint

import (
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"
)

//Service : sprint service struct require repo
type Service struct {
	repo Repository
}

//NewSprintService : create a new service with repo
func NewSprintService(database Repository) *Service {
	return &Service{
		repo: database,
	}
}

//Create : create a new sprint record and insert it to db
func (s *Service) Create(newSprint *Sprint) (*Sprint, error) {

	ID := tools.GenerateStringUUID()
	record, err := NewSprint(newSprint, ID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(record)

	if err != nil {
		return nil, err
	}

	return record, nil
}

//Update : update a new sprint record and insert it to db
func (s *Service) Update(newSprint *Sprint) (*Sprint, error) {
	record, err := UpdateSprint(newSprint)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(record)

	if err != nil {
		return nil, err
	}

	return record, nil
}

//DeleteByID : delete record by id
func (s *Service) DeleteByID(ID string) error {
	if err := s.repo.DeleteByID(ID); err != nil {
		return err
	}
	return nil
}

//FindByID : get record by id
func (s *Service) FindByID(ID string) (*Sprint, error) {
	record, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return record, nil
}
