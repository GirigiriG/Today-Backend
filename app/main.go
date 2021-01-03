package main

import (
	"log"
	"net/http"

	repository "github.com/GirigiriG/Clean-Architecture-golang/pkg/repository/postgres"
	"github.com/GirigiriG/Clean-Architecture-golang/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	db := repository.NewPostgresConnect()

	routes.HandleRoutes(db, router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
