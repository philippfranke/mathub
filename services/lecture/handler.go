package lecture

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	lectures, err := All(mux.Vars(r)["uni"])
	if err != nil {
		return err
	}

	return WriteJSON(w, lectures)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) error {
	lecture, err := Get(mux.Vars(r)["lecture"], mux.Vars(r)["uni"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, lecture)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	var lecture Lecture
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&lecture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	uid, err := strconv.ParseInt(mux.Vars(r)["uni"], 10, 0)
	if err != nil {
		return err
	}

	lecture.UniversityId = uid

	lecture, err = Create(lecture)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, lecture)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	var lecture Lecture
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&lecture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	original, err := Get(mux.Vars(r)["lecture"], mux.Vars(r)["uni"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	lecture.Id = original.Id

	err = Update(lecture)
	return WriteJSON(w, lecture)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) error {
	err := Destroy(mux.Vars(r)["lecture"], mux.Vars(r)["uni"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
