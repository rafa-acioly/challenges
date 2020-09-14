package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rafa-acioly/challenges/entities"

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

// Get represents the HTTP Get verb to get a single walk resource
func (rsc *Resource) Get(w http.ResponseWriter, r *http.Request) {
	showAll, err := strconv.ParseBool(r.URL.Query().Get("show_all"))
	if err != nil {
		showAll = false
	}

	w.Header().Set("content-type", "application/json")
	responseEncoder := json.NewEncoder(w)

	walkings, err := rsc.repository.Index(showAll, 0)
	if err != nil {
		responseEncoder.Encode(NewHTTPError(nil, http.StatusInternalServerError, err))
	}

	responseEncoder.Encode(walkings)
}

// Post represents the HTTP Post verb
func (rsc *Resource) Post(w http.ResponseWriter, r *http.Request) {
	walk := entities.NewWalk()
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)

	requestContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(requestContent, &walk); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = walk.Valid(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(NewHTTPError(err, http.StatusBadRequest, err))
		return
	}

	_, err = rsc.repository.Create(walk)
	if err != nil {
		encoder.Encode(NewHTTPError(nil, http.StatusInternalServerError, err))
		return
	}

	encoder.Encode(walk)
}

// Put represents the HTTP Put verb
func (rsc *Resource) Put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Put resource")
}
