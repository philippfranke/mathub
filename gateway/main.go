// Mathub API-Gateway is a webserver which combines api entrypoints
// into a single one.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
	. "github.com/philippfranke/mathub/shared"
)

var (
	listenAddr = flag.String("listen", ":8080", "HTTP listen address")
	dataPath   = flag.String("data", "./repos", "Repository path")
	dumpPath   = flag.String("dump", "", "mySQL dump path")
	DB         *sqlx.DB
)

func init() {
	flag.Parse()

	var err error
	DB, err = datastore.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if *dumpPath != "" {
		datastore.ImportDump(*dumpPath)
	}
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", Handler(serveRoutes))

	// Register routes
	for _, route := range Routes {
		r.Handle(route.path, route.handler)
	}

	// Start HTTP server
	log.Fatal(http.ListenAndServe(*listenAddr, r))
}

// serveRoutes lists all entrypoints
func serveRoutes(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, Routes)
}

// TODO(franke): StripPrefix in order to avoid duplicated routes
