package user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"crypto/sha256"
	"crypto/subtle"
	"errors"

	"github.com/gorilla/mux"
	. "github.com/philippfranke/mathub/shared"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := All()
	if err != nil {
		return err
	}

	return WriteJSON(w, users)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) error {
	user, err := Get(mux.Vars(r)["user"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return err
	} else if err != nil {
		return err
	}

	return WriteJSON(w, user)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) error {
	var user User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	user, err = Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return WriteJSON(w, user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	var userJson User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&userJson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	userDB, err := GetByEmail(userJson.Email)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return err
	} else if err != nil {
		return err
	}

	if !SecureCompare(userJson.PasswordHash, userDB.PasswordHash) {
		w.WriteHeader(http.StatusForbidden)
		return errors.New("Password is incorrect")
	}

	userDB.PasswordHash = ""

	return WriteJSON(w, userDB)
}

// SecureCompare performs a constant time compare of two strings to limit timing attacks.
func SecureCompare(given string, actual string) bool {
	givenSha := sha256.Sum256([]byte(given))
	actualSha := sha256.Sum256([]byte(actual))

	return subtle.ConstantTimeCompare(givenSha[:], actualSha[:]) == 1
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	var user User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := d.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	original, err := Get(mux.Vars(r)["user"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return err
	} else if err != nil {
		return err
	}

	user.Id = original.Id

	err = Update(user)

	user.PasswordHash = ""

	return WriteJSON(w, user)
}

func DestroyHandler(w http.ResponseWriter, r *http.Request) error {
	err := Destroy(mux.Vars(r)["user"])
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
