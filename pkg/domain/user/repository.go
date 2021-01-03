package user

//Repository database interface
type Repository interface {
	FindByID(string) (*User, error)
	UpdateByID(*User) error
	DeleteByID(string) error
	Create(*User) error
}
