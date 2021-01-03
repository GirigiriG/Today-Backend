package user

import (
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"
)

//Service struct holds repo and user struct
type Service struct {
	repo Repository
}

//NewService register database repo imp
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//Create create new user
func (service *Service) Create(newUser *User) (*User, error) {

	u, err := NewUser(newUser)
	u.ID = tools.GenerateStringUUID()
	if err != nil {
		return nil, err
	}

	err = service.repo.Create(u)

	if err != nil {
		return nil, err
	}
	return u, nil
}

//UpdateByID : Update user by id
func (service *Service) UpdateByID(u *User) (*User, error) {
	record, err := UpdateUser(u)
	if err != nil {
		return nil, err
	}
	err = service.repo.UpdateByID(u)
	if err != nil {
		return nil, err
	}
	return record, nil
}

//FindByID : Get user by ID (string)
func (service *Service) FindByID(ID string) (*User, error) {
	record, err := service.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

//DeleteByID : Delete user by id
func (service *Service) DeleteByID(ID string) error {
	return service.repo.DeleteByID(ID)
}
