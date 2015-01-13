package comment

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/philippfranke/mathub/datastore"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = datastore.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
