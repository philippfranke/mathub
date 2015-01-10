package assignment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/university"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	lectures, err := All(mux.Vars(r)["lecture"])
	if err != nil {
		return err
	}

	return WriteJSON(w, lectures)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	var err error
	var assignment Assignment

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := d.Decode(&assignment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	lectureId := mux.Vars(r)["lecture"]
	universityId := mux.Vars(r)["uni"]

	uni, err := university.Get(universityId)
	if err != nil {
		return err
	}
	lect, err := lecture.Get(lectureId, universityId)
	if err != nil {
		return err
	}

	intLectureId, _ := strconv.ParseInt(lectureId, 10, 0)
	assignment.LectureId = intLectureId

	// Create
	rp := &Repo{uni: uni.Name, lecture: lect.Name}
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
	Update(assignment)
	return WriteJSON(w, assignment)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
