package user

import (
	"errors"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"
	uuid "github.com/satori/go.uuid"
)

//User struct
type User struct {
	ID        uuid.UUID `json: "id"`
	FirstName string    `json: "first_name"`
	LastName  string    `json: "last_name"`
	Email     string    `json: "email"`
}

//CreateNewUser create new user record
func (u *User) CreateNewUser(newUser *User) (*User, error) {

	if newUser.FirstName == "" {

		return nil, errors.New("Please provide first name")
	}

	if newUser.LastName == "" {
		return nil, errors.New("Please provide last name")
	}
	if newUser.Email == "" {
		return nil, errors.New("Please provide email")
	}

	ID := tools.CreateUUID()

	creatdUser := &User{
		ID:        ID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
	}

	return creatdUser, nil

}

func (u *User) validateUserName(name string) error {
	if len(name) < 4 {
		return errors.New("Username must be greater than 4")
	}
	return nil
}
