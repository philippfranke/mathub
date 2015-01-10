package university

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	unis, err := All()
	if err != nil {
		return err
	}

	return WriteJSON(w, unis)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) error {
	uni, err := Get(mux.Vars(r)["uni"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, uni)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	var uni University
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&uni)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	uni, err = Create(uni)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, uni)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	var uni University
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&uni)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	original, err := Get(mux.Vars(r)["uni"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	uni.Id = original.Id

	err = Update(uni)
	return WriteJSON(w, uni)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) error {
	err := Destroy(mux.Vars(r)["uni"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
