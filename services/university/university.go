package university

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

// Router returns entrypoints for university
func Router() http.Handler {
	r := mux.NewRouter()

	r.Handle("/unis", Handler(IndexHandler)).Methods("GET", "HEAD")
	r.Handle("/unis", Handler(CreateHandler)).Methods("POST")
	r.Handle("/unis/{uni}", Handler(ShowHandler)).Methods("GET", "HEAD")
	r.Handle("/unis/{uni}", Handler(UpdateHandler)).Methods("PATCH")
	r.Handle("/unis/{uni}", Handler(DestroyHandler)).Methods("DELETE")

	return r
}

func PreCheck(uni string) bool {
	_, err := Get(uni)
	if err != nil {
		return false
	}

	return true
}
