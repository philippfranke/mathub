package version

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/philippfranke/mathub/services/user"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, a Reference) error {
	assignment, err := All(strconv.FormatInt(a.Id, 10), a.Type)
	if err != nil {
		return err
	}

	return WriteJSON(w, assignment)
}

func ShowHandler(w http.ResponseWriter, r *http.Request, a Reference) error {

	version, err := Get(strconv.FormatInt(a.Id, 10), a.Type, mux.Vars(r)["version"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	var folder string
	if (a.User == user.User{}) {
		folder = filepath.Join(a.University.Name, a.Lecture.Name)
	} else {
		folder = filepath.Join(a.User.Name)
	}

	rp := &Repo{DataPath: DataPath}

	rp.Open(folder)
	tex := rp.ShowFile(strconv.FormatInt(a.Id, 10)+".tex", version.CommitHash)
	version.Tex = tex
	return WriteJSON(w, version)

}

func UpdateHandler(w http.ResponseWriter, r *http.Request, a Reference) error {

	version, err := Get(strconv.FormatInt(a.Id, 10), a.Type, mux.Vars(r)["version"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	var folder string
	if true {
		folder = filepath.Join(a.University.Name, a.Lecture.Name)
	} else {
		folder = filepath.Join(a.User.Name)
	}
	rp := &Repo{DataPath: DataPath}

	rp.Open(folder)
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
