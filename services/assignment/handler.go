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

type Version struct {
	CommitHash    string `db:"commit_hash"`
	ReferenceType string `json:"-,omitempty" db:"ref_type"`
	ReferenceId   int64  `json:"-,omitempty" db:"ref_id"`
	UserId        int64  `db:"user_id"`
	Number        int64  `db:"version"`
	Tex           string `json:"tex,omitempty"`
}

func CreateVersion(v Version) error {
	_, err := DB.Exec("INSERT INTO versions (SELECT ?,?,?,0, IFNULL((max(version)+1),1)  FROM versions where ref_Type=? and ref_id=?)", v.CommitHash, v.ReferenceType, v.ReferenceId, v.ReferenceType, v.ReferenceId)

	if err != nil {
		return err
	}

	return nil
}

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
	v := Version{
		CommitHash:    assignment.CommitHash,
		ReferenceType: "assignments",
		ReferenceId:   assignment.Id,
		UserId:        1,
	}
	err = CreateVersion(v)
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
	if err != nil {
		return err
	}

	v := Version{
		CommitHash:    assignment.CommitHash,
		ReferenceType: "assignments",
		ReferenceId:   assignment.Id,
		UserId:        1,
	}
	err = CreateVersion(v)
	if err != nil {
		return err
	}

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

	out, err := rp.Status()
	if err != nil {
		return err
	}
	if out == "" {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}

	if err := rp.Commit("bla", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
