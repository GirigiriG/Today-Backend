package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/sprint"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/project"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"

	middleware "github.com/GirigiriG/Clean-Architecture-golang/middlerware"

	googleoauth "github.com/GirigiriG/Clean-Architecture-golang/pkg"
	"golang.org/x/oauth2"

	delivery "github.com/GirigiriG/Clean-Architecture-golang/pkg/delivery/http"
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/user"
	routes "github.com/GirigiriG/Clean-Architecture-golang/pkg/routes/unprotected"
	"github.com/gorilla/mux"
)

func HandleRoutes(db *sql.DB, router *mux.Router) {

	//unprotected routes
	unprocted := routes.NewUnprotectedRoutesHandler(router)
	unprocted.InitProtectedRoutes()

	//protected routes

	//user handler/service
	userRepo := user.NewPostgressRepo(db)
	userService := user.NewService(userRepo)
	userRoutesHandler := delivery.NewUserHandler(userService, router)
	userRoutesHandler.HandleRoutes()

	//Task handler/service
	taskRepo := task.NewTaskRepo(db)
	taskService := task.NewTaskService(taskRepo)
	taskeRouterHandler := delivery.NewTaskHandler(taskService, router)
	taskeRouterHandler.HandleRoutes()

	//Project handler/service
	projectRepo := project.NewProjectRepo(db)
	projectService := project.NewProjectService(projectRepo)
	projectRouterHandler := delivery.NewProjectHandler(projectService, router)
	projectRouterHandler.HandleRoutes()

	//Sprint handler/service

	sprintRepo := sprint.NewSprintRepositroy(db)
	sprintService := sprint.NewSprintService(sprintRepo)
	sprintRouterHandler := delivery.NewSprintHandler(sprintService, router)
	sprintRouterHandler.HandleRoutes()

	router.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.FormValue("state") != googleoauth.RandomState {
			fmt.Println("State is not valid")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		token, err := googleoauth.GoogleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))

		if !token.Valid() {
			fmt.Printf("Could not obtain token: %s", err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
		if err != nil {
			fmt.Printf("Get request failed: %s", err.Error())

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Printf("Unable to parse response: %s", err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		var googleResp *googleoauth.GoogleAuthResponse

		json.Unmarshal(content, &googleResp)
		googleResp.Token = middleware.GenerateJWToken()
		json.NewEncoder(w).Encode(googleResp)

		result, err := json.Marshal(googleResp)
		if err != nil {
			fmt.Printf("Unable to unmarshal json: %s", err.Error())
			return
		}
		json.NewEncoder(w).Encode(result)
	})

}
