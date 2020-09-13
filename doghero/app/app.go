package app

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rafa-acioly/challenges/config"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// App represents a basic app structure
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize set the all resources in order to run the application,
// it will set the database connection and the http router
func (a App) Initialize(settings config.Config)  {
	a.SetDatabaseConnection(settings)
	a.Router = mux.NewRouter()
}

// SetDatabaseConnection creates a database connection
func (a App) SetDatabaseConnection(settings config.Config)  {
	database, err := sql.Open("sqlite3", "../doghero.db")
	if err != nil {
		log.Fatal("Database connection issue: " + err.Error())
	}

	a.DB = database
}

// Run starts a http server
func (a App) Run(port string) {
	defer a.DB.Close()

	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}