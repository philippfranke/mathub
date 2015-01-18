package shared

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler represents http.Handler with a proper error handling
type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := h(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

// WriteJSON encodes an interface and sets the content-type
func WriteJSON(w http.ResponseWriter, i interface{}) error {
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, OPTIONS")
	w.Write(b)

	return nil
}
