package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/user"
	repository "github.com/GirigiriG/Clean-Architecture-golang/pkg/repository/postgres"
	"github.com/gorilla/mux"
)

func TestUserHandlerFind(t *testing.T) {
	endPoint := "localhost:3001/user/find/89e8a24d-09cc-4ce6-b9b6-cc2908ab1542222"
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	recorder := httptest.NewRecorder()
	handler := createUserHandler()

	http.HandlerFunc(handler.FindByID).ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code, "expected 200")
}

func TestUserHandlerCreate(t *testing.T) {
	content := `{
		"FirstName": "Gideon",
		"LastName": "Girigiri",
		"Email": "me@gmail.com"
	}`

	endPoint := "localhost:3001/user/create"
	req, err := http.NewRequest("GET", endPoint, strings.NewReader(content))
	if err != nil {
		t.Errorf(err.Error())
	}
	recorder := httptest.NewRecorder()
	handler := createUserHandler()

	http.HandlerFunc(handler.FindByID).ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code, "expected 200")
}

func createUserHandler() *UserHandler {
	router := mux.NewRouter()
	db := repository.NewPostgresConnect()

	userRepo := user.NewPostgressRepo(db)
	userService := user.NewService(userRepo)
	userRoutesHandler := NewUserHandler(userService, router)
	userRoutesHandler.HandleRoutes()

	return userRoutesHandler
}
