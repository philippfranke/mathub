package lecture

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
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

// Router returns entrypoints for lecture
func Router() http.Handler {

	r := mux.NewRouter()

	r.Handle("/unis/{uni}/lectures", Handler(filterHandler(IndexHandler))).Methods("GET", "HEAD")
	r.Handle("/unis/{uni}/lectures", Handler(filterHandler(CreateHandler))).Methods("POST")
	r.Handle("/unis/{uni}/lectures/{lecture}", Handler(filterHandler(ShowHandler))).Methods("GET", "HEAD")
	r.Handle("/unis/{uni}/lectures/{lecture}", Handler(filterHandler(UpdateHandler))).Methods("PATCH")
	r.Handle("/unis/{uni}/lectures/{lecture}", Handler(filterHandler(DestroyHandler))).Methods("DELETE")

	return r
}

// Filter valid uniID
type UniversityHandler func(w http.ResponseWriter, r *http.Request, u university.University) error

func filterHandler(next UniversityHandler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		u, err := university.Get(mux.Vars(r)["uni"])
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return nil
		} else if err != nil {
			return err
		}

		return next(w, r, u)
	}
}
