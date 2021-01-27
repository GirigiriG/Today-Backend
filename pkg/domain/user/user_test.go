package user_test

import (
	"testing"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	u := createTestData()
	newUser, err := user.NewUser(u)

	assert.Equal(t, nil, err)
	assert.NotNil(t, newUser)
}

func TestFirstnameEmpty(t *testing.T) {
	u := createTestData()
	u.FirstName = ""
	newUser, err := user.NewUser(u)

	assert.Equal(t, user.ErrFirstNameRequired, err.Error())
	assert.Nil(t, newUser)
}

func TestLastnameEmpty(t *testing.T) {
	u := createTestData()
	u.LastName = ""
	newUser, err := user.NewUser(u)

	assert.Equal(t, user.ErrLastNameRequired, err.Error())
	assert.Nil(t, newUser)
}

func TestEmailEmpty(t *testing.T) {
	u := createTestData()
	u.Email = ""

	newUser, err := user.NewUser(u)

	assert.Equal(t, user.ErrEmailRequired, err.Error())
	assert.Nil(t, newUser)
}

func createTestData() *user.User {
	return &user.User{
		ID:        "12345",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com.invalid",
	}
}
