package solution

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
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

// Router returns entrypoints for lecture
func Router(path string) http.Handler {
	DataPath, _ = filepath.Abs(path)

	r := mux.NewRouter()

	r.Handle("/users/{user}/solutions", Handler(filterHandler(IndexHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/users/{user}/solutions/{solution}", Handler(filterHandler(ShowHandler))).Methods("GET", "HEAD", "OPTIONS")
	r.Handle("/users/{user}/solutions/{solution}", Handler(filterHandler(UpdateHandler))).Methods("PATCH")
	r.Handle("/users/{user}/solutions/{solution}", Handler(filterHandler(DestroyHandler))).Methods("DELETE")
	r.Handle("/users/{user}/solutions", Handler(filterHandler(CreateHandler))).Methods("POST")

	return r
}

type FilterHandler func(w http.ResponseWriter, r *http.Request, u user.User) error

func filterHandler(next FilterHandler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := user.Get(mux.Vars(r)["user"])
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return nil
		} else if err != nil {
			return err
		}

		return next(w, r, u)
	}
}
