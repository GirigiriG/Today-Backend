package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/auth"
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/sprint"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/project"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"

	"golang.org/x/oauth2"

	delivery "github.com/GirigiriG/Clean-Architecture-golang/pkg/delivery/http"
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/user"
	routes "github.com/GirigiriG/Clean-Architecture-golang/pkg/routes/unprotected"
	"github.com/gorilla/mux"
)

//HandleRoutes all entity routes
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
		if r.FormValue("state") != auth.RandomState {
			fmt.Println("State is not valid")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		token, err := auth.GoogleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))

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

		var googleResp *auth.GoogleAuthResponse

		json.Unmarshal(content, &googleResp)
		googleResp.Token = auth.GenerateJWToken()
		json.NewEncoder(w).Encode(googleResp)

		result, err := json.Marshal(googleResp)
		if err != nil {
			fmt.Printf("Unable to unmarshal json: %s", err.Error())
			return
		}
		json.NewEncoder(w).Encode(result)
	})

	router.HandleFunc("/get/record/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		type record struct {
			ID   string
			Name string
		}

		var results []record
		var typeRecord record

		name := r.URL.Query()["name"][0]

		if len(name) == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(results)
			return
		}
		recordType := r.URL.Query()["type"][0]

		query := `SELECT id, name FROM ` + recordType + ` WHERE name ILIKE ` + "'%" + name + "%' " + `LIMIT 5;`
		query = strings.ReplaceAll(query, "+", " ")

		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			rows.Scan(&typeRecord.ID, &typeRecord.Name)
			results = append(results, typeRecord)
		}

		if len(results) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return

		}
		json.NewEncoder(w).Encode(results)

	})
}
