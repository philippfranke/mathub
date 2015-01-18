package comment

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	"github.com/philippfranke/mathub/services/assignment"
	. "github.com/philippfranke/mathub/shared"
	//TODO: import Solution
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = datastore.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

// Router returns entrypoints for comments
func Router() http.Handler {
	r := mux.NewRouter()

	r.Handle("/comments/{refType}/{refId}", FilterHandler(Handler(IndexHandler))).Methods("GET", "HEAD")
	r.Handle("/comments", Handler(CreateHandler)).Methods("POST")
	r.Handle("/comments/{comment}", Handler(ShowHandler)).Methods("GET", "HEAD")
	r.Handle("/comments/{comment}", Handler(UpdateHandler)).Methods("PATCH")
	r.Handle("/comments/{comment}", Handler(DestroyHandler)).Methods("DELETE")

	return r
}

func PreCheck(refType, refId string) bool {
	switch refType {
	case "assignment":
		_, err := assignment.Get(refId)
		if err != nil {
			return false
		}
	case "solution": //TODO:
		/*
			_, err := solution.Get(refId)
			if err != nil {
				return false
			}
		*/
	default:
		return false
	}

	return true
}

// Filter valid refId
type FilterHandler func(w http.ResponseWriter, r *http.Request) error

func (h FilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !PreCheck(mux.Vars(r)["refType"], mux.Vars(r)["refId"]) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h(w, r)

}
