package lecture

import (
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

	r.Handle("/unis/{uni}/lectures", FilterHandler(Handler(IndexHandler))).Methods("GET", "HEAD")
	r.Handle("/unis/{uni}/lectures", FilterHandler(Handler(CreateHandler))).Methods("POST")
	r.Handle("/unis/{uni}/lectures/{lecture}", FilterHandler(Handler(ShowHandler))).Methods("GET", "HEAD")
	r.Handle("/unis/{uni}/lectures/{lecture}", FilterHandler(Handler(UpdateHandler))).Methods("PATCH")
	r.Handle("/unis/{uni}/lectures/{lecture}", FilterHandler(Handler(DestroyHandler))).Methods("DELETE")

	return r
}

// Filter valid uniID
type FilterHandler func(w http.ResponseWriter, r *http.Request) error

func (h FilterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !university.PreCheck(mux.Vars(r)["uni"]) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h(w, r)

}
