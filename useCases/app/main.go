package main

import (
	"log"
	"net/http"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/routes"

	repository "github.com/GirigiriG/Clean-Architecture-golang/pkg/repository/postgres"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()

	db := repository.NewPostgresConnect()

	routes.HandleRoutes(db, router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
