package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rafa-acioly/challenges/repositories"
)

// Resource represents the api handler resource
type Resource struct {
	repository repositories.Repository
}

// NewDogWalkResource retrieves a new instance of the resource with http handlers
func NewDogWalkResource(db *sql.DB) Resource {
	return Resource{repository: repositories.NewDogWalkRepository(db)}
}

// Routes attach the http verbs and handlers to the resource
func (rsc *Resource) Routes(router chi.Router) {
	router.Get("/", rsc.Get)
	router.Post("/", rsc.Post)
	router.Put("/{id}/", rsc.Put)
}

// Get represents the HTTP Get verb
func (rsc *Resource) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get resource")
}

// Post represents the HTTP Post verb
func (rsc *Resource) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post resource")
}

// Put represents the HTTP Put verb
func (rsc *Resource) Put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Put resource")
}
