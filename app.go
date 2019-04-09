package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App stores application details
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize will create database connection and wire up routes
func (app *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		user,
		password,
		dbname,
		"disable",
	)

	fmt.Println(connectionString)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
}

// Run will start the application
func (app *App) Run(addr string) {}
