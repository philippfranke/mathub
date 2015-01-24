package version

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	"github.com/philippfranke/mathub/services/assignment"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/university"

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

func Router(path string) http.Handler {
	r := mux.NewRouter()

	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}/versions", Handler(filterHandler(IndexHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}/versions/{version}", Handler(filterHandler(ShowHandler))).Methods("GET")
	r.Handle("/unis/{uni}/lectures/{lecture}/assignments/{assignment}/versions/{version}", Handler(filterHandler(UpdateHandler))).Methods("PATCH")

	return r
}

type FilterHandler func(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture, a assignment.Assignment) error

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

		a, err := assignment.Get(mux.Vars(r)["assignment"])
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return nil
		} else if err != nil {
			return err
		}

		return next(w, r, u, l, a)
	}
}
