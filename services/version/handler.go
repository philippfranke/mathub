package version

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philippfranke/mathub/services/assignment"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/university"

	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture, a assignment.Assignment) error {
	assignment, err := All(strconv.FormatInt(a.Id, 10), "assignments")
	if err != nil {
		return err
	}

	return WriteJSON(w, assignment)
}

func ShowHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture, a assignment.Assignment) error {

	version, err := Get(strconv.FormatInt(a.Id, 10), "assignments", mux.Vars(r)["version"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	rp := assignment.NewRepo(u.Name, l.Name)

	rp.Open()
	tex := rp.ShowFile(strconv.FormatInt(a.Id, 10)+".tex", version.CommitHash)
	version.Tex = tex
	return WriteJSON(w, version)

}

func UpdateHandler(w http.ResponseWriter, r *http.Request, u university.University, l lecture.Lecture, a assignment.Assignment) error {

	version, err := Get(strconv.FormatInt(a.Id, 10), "assignments", mux.Vars(r)["version"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	rp := assignment.NewRepo(u.Name, l.Name)

	rp.Open()
	tex := rp.ShowFile(strconv.FormatInt(a.Id, 10)+".tex", version.CommitHash)

	err = rp.Update(strconv.FormatInt(a.Id, 10)+".tex", tex)
	if err != nil {
		return err
	}

	out, err := rp.Status()
	if err != nil {
		return err
	}
	if out == "" {
		return WriteJSON(w, version)
	}

	if err := rp.Commit("Revert", "Tim Trompete <mail@mail.com>"); err != nil {
		return err
	}

	version.CommitHash = rp.LastHash()

	Create(version)

	return WriteJSON(w, version)
}
