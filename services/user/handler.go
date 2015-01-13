package user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := All()
	if err != nil {
		return err
	}

	return WriteJSON(w, users)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) error {
	user, err := Get(mux.Vars(r)["user"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, user)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	var user User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	user, err = Create(user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, user)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	var user User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	original, err := Get(mux.Vars(r)["user"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	user.Id = original.Id

	err = Update(user)
	return WriteJSON(w, user)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) error {
	err := Destroy(mux.Vars(r)["user"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
