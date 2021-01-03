package user_test

import (
	"testing"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

func Test_createNewUser(t *testing.T) {
	u := createTestData()
	newUser, err := user.NewUser(u)

	assert.Equal(t, nil, err)
	assert.NotNil(t, newUser)
}

func Test_Firstname_Empty(t *testing.T) {
	u := createTestData()
	u.FirstName = ""
	newUser, err := user.NewUser(u)

	assert.Equal(t, user.FirstNameRequired, err.Error())
	assert.Nil(t, newUser)
}

func Test_Lastname_Empty(t *testing.T) {
	u := createTestData()
	u.LastName = ""
	newUser, err := user.NewUser(u)

	assert.Equal(t, user.LastNameRequired, err.Error())
	assert.Nil(t, newUser)
}

func Test_Email_Empty(t *testing.T) {
	u := createTestData()
	u.Email = ""

	newUser, err := user.NewUser(u)

	assert.Equal(t, user.EmailRequired, err.Error())
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
