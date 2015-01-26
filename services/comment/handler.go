package comment

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	commentTree, err := All(mux.Vars(r)["refType"], mux.Vars(r)["refId"])
	if err != nil {
		return err
	}

	return WriteJSON(w, commentTree)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) error {
	comment, err := Get(mux.Vars(r)["comment"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, comment)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	var comment Comment
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	comment, err = Create(comment)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, comment)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	var comment Comment
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	original, err := Get(mux.Vars(r)["comment"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	comment.Id = original.Id

	err = Update(comment)
	if err != nil {
		return err
	}
	return WriteJSON(w, comment)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) error {
	err := Destroy(mux.Vars(r)["comment"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
