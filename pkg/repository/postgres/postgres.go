package repository

import (
	"database/sql"
	"fmt"
	"os"
)

//NewPostgresConnect connect to Pg database
func NewPostgresConnect() *sql.DB {
	const (
		host = "localhost"
		port = 5432
	)

	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	connectionString := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	connStr := fmt.Sprintf(connectionString, host, port, userName, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}
	
	return db
}
