package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	// Initialize empty string
	var user, password, host, port, dbname string

	// Check app env for db connection credentials
	switch APP_ENV := os.Getenv("APP_ENV"); APP_ENV {
	case "development":
		user = os.Getenv("DB_USER_DEV")
		password = os.Getenv("DB_PASS_DEV")
		host = os.Getenv("DB_HOST_DEV")
		port = os.Getenv("DB_PORT_DEV")
		dbname = os.Getenv("DB_NAME_DEV")
	case "test":
		user = os.Getenv("DB_USER_TEST")
		password = os.Getenv("DB_PASS_TEST")
		host = os.Getenv("DB_HOST_TEST")
		port = os.Getenv("DB_PORT_TEST")
		dbname = os.Getenv("DB_NAME_TEST")
	case "production":
		user = os.Getenv("DB_USER_PROD")
		password = os.Getenv("DB_PASS_PROD")
		host = os.Getenv("DB_HOST_PROD")
		port = os.Getenv("DB_PORT_PROD")
		dbname = os.Getenv("DB_NAME_PROD")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(port string) {}
