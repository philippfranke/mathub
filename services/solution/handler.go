package solution

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philippfranke/mathub/services/user"
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

func IndexHandler(w http.ResponseWriter, r *http.Request, u user.User) error {
	solution, err := All(strconv.FormatInt(u.Id, 10))
	if err != nil {
		return err
	}

	return WriteJSON(w, solution)
}

func ShowHandler(w http.ResponseWriter, r *http.Request, u user.User) error {
	solution, err := Get(mux.Vars(r)["solution"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	return WriteJSON(w, solution)
}

func CreateHandler(w http.ResponseWriter, r *http.Request, u user.User) error {
	var err error
	var solution Solution

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := d.Decode(&solution); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	solution.UserId = u.Id

	// Create

	rp := &Repo{DataPath: DataPath}
	folder := filepath.Join(u.Name)
	if err := rp.Create(folder); err != nil {
		return err
	}

	solution, err = Create(solution)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%d.tex", solution.Id)
	if err := rp.Add(filename, solution.Tex); err != nil {
		return err
	}

	if err := rp.Commit("Initial commit", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}

	solution.CommitHash = rp.LastHash()
	UpdateId(solution)
	v := Version{
		CommitHash:    solution.CommitHash,
		ReferenceType: "solutions",
		ReferenceId:   solution.Id,
		UserId:        1,
	}
	err = CreateVersion(v)
	return WriteJSON(w, solution)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, u user.User) error {

	var solution Solution

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := d.Decode(&solution); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	assignmentId := mux.Vars(r)["solution"]

	original, err := Get(assignmentId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	if solution.UserId == 0 {
		solution.UserId = u.Id
	}

	// Open
	folder := filepath.Join(u.Name)
	rp := &Repo{DataPath: DataPath}
	rp.Open(folder)
	filename := fmt.Sprintf("%s.tex", assignmentId)
	if err := rp.Update(filename, solution.Tex); err != nil {
		return err
	}

	solution.Id = original.Id

	out, err := rp.Status()
	if err != nil {
		return err
	}
	if out == "" {
		return WriteJSON(w, solution)
	}

	if err := rp.Commit("bla", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}
	solution.CommitHash = rp.LastHash()

	err = Update(solution)
	if err != nil {
		return err
	}

	v := Version{
		CommitHash:    solution.CommitHash,
		ReferenceType: "solutions",
		ReferenceId:   solution.Id,
		UserId:        1,
	}
	err = CreateVersion(v)
	if err != nil {
		return err
	}

	return WriteJSON(w, solution)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request, u user.User) error {

	solutionId := mux.Vars(r)["solution"]

	folder := filepath.Join(u.Name)
	rp := &Repo{DataPath: DataPath}
	rp.Open(folder)

	filename := fmt.Sprintf("%s.tex", solutionId)
	rp.Destroy(filename)

	err := Destroy(mux.Vars(r)["solution"])
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
