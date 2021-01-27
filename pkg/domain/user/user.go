package user

import (
	"errors"
)

//ErrFirstNameRequired : first name required message
var ErrFirstNameRequired = errors.New("first name is required")

//ErrLastNameRequired : last name required message
var ErrLastNameRequired = errors.New("Last name is required")

//ErrEmailRequired : email required message
var ErrEmailRequired = errors.New("Email address is required")

//User struct
type User struct {
	ID        string
	Name      string
	FirstName string
	LastName  string
	Email     string
}

//NewUser create new user record
func NewUser(u *User) (*User, error) {
	if err := validate(u); err != nil {
		return nil, err
	}
	u.Name = u.FirstName + " " + u.LastName
	return u, nil
}

//UpdateUser : update the user record.
func UpdateUser(u *User) (*User, error) {
	if err := validate(u); err != nil {
		return nil, err
	}
	return u, nil
}

func validate(u *User) error {
	if len(u.FirstName) == 0 {
		return ErrFirstNameRequired
	}

	if len(u.LastName) == 0 {
		return ErrLastNameRequired
	}
	if len(u.Email) == 0 {
		return ErrEmailRequired
	}
	return nil
}
