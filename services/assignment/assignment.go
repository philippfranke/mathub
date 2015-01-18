package assignment

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/university"
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

	r.Handle("/unis/{uni}/lectures/{lecture}/assignments", Handler(filterHandler(IndexHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}", Handler(filterHandler(ShowHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}", Handler(filterHandler(UpdateHandler))).Methods("PATCH")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}", Handler(filterHandler(DestroyHandler))).Methods("DELETE")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments", Handler(filterHandler(CreateHandler))).Methods("POST")

	return r
}

type FilterHandler func(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture) error

func filterHandler(next FilterHandler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := university.Get(mux.Vars(r)["uni"])
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return nil
		} else if err != nil {
			return err
		}

		l, err := lecture.Get(mux.Vars(r)["lecture"])
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return nil
		} else if err != nil {
			return err
		}

		return next(w, r, u, l)
	}
}
