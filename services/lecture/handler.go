package lecture

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philippfranke/mathub/services/university"
	"github.com/philippfranke/mathub/services/user"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, u university.University) error {
	lectures, err := All(strconv.FormatInt(u.Id, 10))
	if err != nil {
		return err
	}

	return WriteJSON(w, lectures)
}

func ShowHandler(w http.ResponseWriter, r *http.Request, u university.University) error {
	lecture, err := Get(mux.Vars(r)["lecture"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, lecture)
}

func CreateHandler(w http.ResponseWriter, r *http.Request, u university.University) error {
	_, err := user.Get(r.Header["User"][0])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	var lecture Lecture
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = d.Decode(&lecture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	lecture.UniversityId = u.Id

	lecture, err = Create(lecture)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, lecture)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, u university.University) error {
	_, err := user.Get(r.Header["User"][0])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	var lecture Lecture
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = d.Decode(&lecture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	original, err := Get(mux.Vars(r)["lecture"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	lecture.Id = original.Id

	if lecture.UniversityId == 0 {
		lecture.UniversityId = original.UniversityId
	}

	err = Update(lecture)
	return WriteJSON(w, lecture)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request, u university.University) error {
	_, err := user.Get(r.Header["User"][0])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	err = Destroy(mux.Vars(r)["lecture"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
