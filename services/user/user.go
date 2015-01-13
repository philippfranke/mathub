package user

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

// Router returns entrypoints for user
func Router() http.Handler {

	r := mux.NewRouter()

	r.Handle("/users", Handler(IndexHandler)).Methods("GET", "HEAD")
	r.Handle("/users", Handler(CreateHandler)).Methods("POST")
	r.Handle("/users/{user}", Handler(ShowHandler)).Methods("GET", "HEAD")
	r.Handle("/users/{user}", Handler(UpdateHandler)).Methods("PATCH")
	r.Handle("/users/{user}", Handler(DestroyHandler)).Methods("DELETE")

	return r
}
