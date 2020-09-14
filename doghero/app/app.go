package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rafa-acioly/challenges/config"

	_ "github.com/lib/pq" // postgres golang driver
)

// App represents a basic app structure
type App struct {
	Router *chi.Mux
	DB     *sql.DB
}

// Initialize set the all resources in order to run the application,
// it will set the database connection and the http router
func (a *App) Initialize(settings config.Config) {
	a.SetDatabaseConnection(settings)
	a.Router = chi.NewRouter()
}

// SetDatabaseConnection creates a database connection
func (a *App) SetDatabaseConnection(settings config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		settings.DB.Host, settings.DB.Port,
		settings.DB.User, settings.DB.Name,
		settings.DB.Password,
	)
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("database connection issue: " + err.Error())
	}

	if err := database.Ping(); err != nil {
		log.Fatal("database connection issue: " + err.Error())
	}

	a.DB = database
}

// AddRoute attach a new route group the the HTTP server
func (a *App) AddRoute(pattern string, fn func(r chi.Router)) {
	a.Router.Route(pattern, fn)
}

// SetMiddlewares defines the middleware used on every http request
func (a *App) SetMiddlewares() {
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.AllowContentType("application/json"))
	a.Router.Use(middleware.Recoverer)
}

// Run starts a http server
func (a *App) Run(port string) {
	defer a.DB.Close()

	fmt.Println("Listening on port: " + port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}
