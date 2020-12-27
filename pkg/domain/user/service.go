package user

//Service struct holds repo and user struct
type Service struct {
	repo Repository
	u    User
}

//NewService register database repo imp
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

//GetUserByID get user by uuid (string)
func (s *Service) GetUserByID(uuid string) (*User, error) {
	return s.repo.FindByID(uuid)
}

//DeleteUserByID delete user by id
func (s *Service) DeleteUserByID(uuid string) error {
	return s.repo.DeleteUserByID(uuid)
}

//CreateNewUser create new user
func (s *Service) CreateNewUser(newUser *User) (*User, error) {

	u, err := s.u.CreateNewUser(newUser)
	if err != nil {
		return nil, err
	}
	
	err = s.repo.Create(u)

	if err != nil {
		return nil, err
	}
	return u, nil
}
