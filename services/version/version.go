package version

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	"github.com/philippfranke/mathub/services/assignment"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/solution"
	"github.com/philippfranke/mathub/services/university"
	"github.com/philippfranke/mathub/services/user"

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

func Router(path string) http.Handler {
	DataPath, _ = filepath.Abs(path)

	r := mux.NewRouter()

	r.Handle("/unis/{uni}/lectures/{lecture}/{ref_type}/{ref_id}/versions", Handler(filterHandler(IndexHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/unis/{uni}/lectures/{lecture}/{ref_type}/{ref_id}/versions/{version}", Handler(filterHandler(ShowHandler))).Methods("GET", "OPTIONS")
	r.Handle("/unis/{uni}/lectures/{lecture}/{ref_type}/{ref_id}/versions/{version}", Handler(filterHandler(UpdateHandler))).Methods("PATCH")
	r.Handle("/users/{user}/{ref_type}/{ref_id}/versions", Handler(filterHandler(IndexHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/users/{user}/{ref_type}/{ref_id}/versions/{version}", Handler(filterHandler(ShowHandler))).Methods("GET", "OPTIONS")
	r.Handle("/users/{user}/{ref_type}/{ref_id}/versions/{version}", Handler(filterHandler(UpdateHandler))).Methods("PATCH")

	return r
}

type FilterHandler func(w http.ResponseWriter, r *http.Request, a Reference) error

func filterHandler(next FilterHandler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {

		var a Reference
		switch mux.Vars(r)["ref_type"] {
		case "assignments":
			u, err := university.Get(mux.Vars(r)["uni"])
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return nil
			} else if err != nil {
				return err
			}
			fmt.Println("bla1")
			l, err := lecture.Get(mux.Vars(r)["lecture"])
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return nil
			} else if err != nil {
				return err
			}
			fmt.Println("bla1")
			b, err := assignment.Get(mux.Vars(r)["ref_id"])
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return nil
			} else if err != nil {
				return err
			}
			a.Id = b.Id
			a.Type = "assignments"
			a.Lecture = l
			a.University = u
		case "solutions":
			u, err := user.Get(mux.Vars(r)["user"])
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return nil
			} else if err != nil {
				return err
			}
			b, err := solution.Get(mux.Vars(r)["ref_id"])
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return nil
			} else if err != nil {
				return err
			}
			a.Id = b.Id
			a.Type = "solutions"
			a.User = u
		default:
			w.WriteHeader(http.StatusNotFound)
			return nil
		}

		return next(w, r, a)
	}
}

type Reference struct {
	Type       string
	Id         int64
	Lecture    lecture.Lecture
	University university.University
	User       user.User
}
