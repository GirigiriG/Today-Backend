package sprint

//Repository : sprint repo interface
type Repository interface {
	Create(*Sprint) error
	Update(*Sprint) error
	DeleteByID(string) error
	FindByID(string) (*Sprint, error)
}
