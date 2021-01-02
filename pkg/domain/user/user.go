package user

import (
	"errors"
)

const (
	FirstNameRequired = "First name is required."
	LastNameRequired  = "Last name is required."
	EmailRequired     = "Email address is required."
)

//User struct
type User struct {
	ID        string `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
}

//NewUser create new user record
func NewUser(newUser *User) (*User, error) {
	if err := validateUserName(newUser); err != nil {
		return nil, err
	}

	creatdUser := &User{
		ID:        newUser.ID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}

	return creatdUser, nil
}

func validateUserName(u *User) error {
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
