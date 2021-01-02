package user

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

//GetUserByID get user by ID (string)
func (s *Service) GetUserByID(ID string) (*User, error) {
	return s.repo.FindByID(ID)
}

//DeleteUserByID delete user by id
func (s *Service) DeleteUserByID(ID string) error {
	return s.repo.DeleteUserByID(ID)
}

//CreateNewUser create new user
func (s *Service) CreateNewUser(newUser *User) (*User, error) {
	
	u, err := NewUser(newUser)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(u)

	if err != nil {
		return nil, err
	}
	return u, nil
}
