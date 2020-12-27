package user

//Repository database interface
type Repository interface {
	FindByID(string) (*User, error)
	UpdateUserByID(string) error
	DeleteUserByID(string) error
	Create(*User) error
}
