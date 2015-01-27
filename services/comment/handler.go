package comment

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/philippfranke/mathub/services/user"
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
	user, err := user.Get(r.Header["User"][0])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	var comment Comment
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = d.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	comment.UserID = user.Id

	comment, err = Create(comment)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, comment)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	user, err := user.Get(r.Header["User"][0])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	var comment Comment
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = d.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	comment.UserID = user.Id

	original, err := Get(mux.Vars(r)["comment"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return nil
	} else if err != nil {
		return err
	}

	comment.Id = original.Id

	commentTypes := reflect.TypeOf(Comment{})
	updateValues := reflect.ValueOf(&comment)

	for i := 0; i < commentTypes.NumField(); i++ {
		var comp = reflect.New(commentTypes.Field(i).Type).Elem().Interface()

		if updateValues.Elem().Field(i).Interface() == comp {
			val := reflect.ValueOf(&original).Elem().Field(i)
			updateValues.Elem().Field(i).Set(val)
		}
	}

	err = Update(comment)
	if err != nil {
		return err
	}
	return WriteJSON(w, comment)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) error {
	_, err := user.Get(r.Header["User"][0])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	err = Destroy(mux.Vars(r)["comment"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
