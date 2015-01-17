package assignment

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/university"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture) error {
	assignment, err := All(strconv.FormatInt(l.Id, 10))
	if err != nil {
		return err
	}

	return WriteJSON(w, assignment)
}

func ShowHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture) error {
	assignment, err := Get(mux.Vars(r)["assignment"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, assignment)
}

func CreateHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture) error {
	var err error
	var assignment Assignment

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := d.Decode(&assignment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	assignment.LectureId = l.Id

	// Create
	rp := &Repo{uni: u.Name, lecture: l.Name}
	if err := rp.Create(); err != nil {
		return err
	}

	assignment, err = Create(assignment)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%d.tex", assignment.Id)
	if err := rp.Add(filename, assignment.Tex); err != nil {
		return err
	}

	if err := rp.Commit("Initial commit", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}

	assignment.CommitHash = rp.LastHash()
	UpdateId(assignment)
	return WriteJSON(w, assignment)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture) error {

	var assignment Assignment

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := d.Decode(&assignment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	assignmentId := mux.Vars(r)["assignment"]

	original, err := Get(assignmentId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	if assignment.LectureId == 0 {
		assignment.LectureId = l.Id
	}

	// Open
	rp := &Repo{uni: u.Name, lecture: l.Name}
	rp.Open()
	filename := fmt.Sprintf("%s.tex", assignmentId)
	if err := rp.Update(filename, assignment.Tex); err != nil {
		return err
	}

	assignment.Id = original.Id

	out, err := rp.Status()
	if err != nil {
		return err
	}
	if out == "" {
		return WriteJSON(w, assignment)
	}

	if err := rp.Commit("bla", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}
	assignment.CommitHash = rp.LastHash()

	err = Update(assignment)

	return WriteJSON(w, assignment)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture) error {

	assignmentId := mux.Vars(r)["assignment"]

	rp := &Repo{uni: u.Name, lecture: l.Name}
	rp.Open()

	filename := fmt.Sprintf("%s.tex", assignmentId)
	rp.Destroy(filename)

	err := Destroy(mux.Vars(r)["assignment"])
	if err != nil {
		return err
	}

	if err := rp.Commit("bla", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
