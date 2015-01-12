package assignment

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	. "github.com/philippfranke/mathub/shared"
)

var DB *sqlx.DB
var DataPath string

func init() {
	var err error
	DB, err = datastore.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

// Router returns entrypoints for lecture
func Router(path string) http.Handler {
	DataPath, _ = filepath.Abs(path)

	r := mux.NewRouter()

	r.Handle("/unis/{uni}/lectures/{lecture}/assignments", Handler(IndexHandler)).Methods("GET", "HEAD")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}", Handler(ShowHandler)).Methods("GET", "HEAD")

	r.Handle("/unis/{uni}/lectures/{lecture}/assignments", Handler(CreateHandler)).Methods("POST")

	return r
}
