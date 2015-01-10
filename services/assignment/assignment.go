package assignment

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	. "github.com/philippfranke/mathub/shared"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = datastore.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

// Router returns entrypoints for lecture
func Router() http.Handler {

	r := mux.NewRouter()

	r.Handle("/unis/{uni}/lectures/{lecture}/assignments", Handler(IndexHandler)).Methods("GET", "HEAD")

	r.Handle("/unis/{uni}/lectures/{lecture}/assignments", Handler(CreateHandler)).Methods("POST")

	return r
}
