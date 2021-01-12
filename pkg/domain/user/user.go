package user

import (
	"errors"
)

const (
	//FirstNameRequired : first name required message
	FirstNameRequired = "First name is required."
	//LastNameRequired : last name required message
	LastNameRequired = "Last name is required."
	//EmailRequired : email required message
	EmailRequired = "Email address is required."
)

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
	u.Name = u.FirstName +" "+ u.LastName
	return u, nil
}

func UpdateUser(u *User) (*User, error) {
	if err := validate(u); err != nil {
		return nil, err
	}
	return u, nil
}

func validate(u *User) error {
	if len(u.FirstName) == 0 {
		return errors.New(FirstNameRequired)
	}

	if len(u.LastName) == 0 {
		return errors.New(LastNameRequired)
	}
	if len(u.Email) == 0 {
		return errors.New(EmailRequired)
	}
	return nil
}
