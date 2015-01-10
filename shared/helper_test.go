package shared

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder() // ResponseWriter

	// Tests invalid type
	f := func(a string) string { return a }
	err := WriteJSON(w, f)
	if err == nil {
		t.Fatalf("writeJSON: %v", err)
	}
	w.Flush()
	// Tests valid type

	input := map[string]int{
		"A": 1,
		"B": 2,
	}

	wantBody := `{
    "A": 1,
    "B": 2
}`

	err = WriteJSON(w, input)
	if err != nil {
		t.Fatalf("writeJSON: %v", err)
	}

	gotHeader := w.Header().Get("Content-Type")
	gotBody := w.Body.String()

	if gotHeader != "application/json; charset=utf-8" {
		t.Errorf("writeJSON: Content-Type is wrong: %s", gotHeader)
	}

	if gotBody != wantBody {
		t.Errorf("writeJSON: body is %s, expected %s", gotBody, wantBody)
	}

}

func TestServeHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Handler: no request %v", err)
	}

	// Tests error
	inputErr := func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("Something went wrong!")
	}

	Handler(inputErr).ServeHTTP(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Handler: %v, expected 500", w.Code)
	}

	// w.Flush()
	w = httptest.NewRecorder()

	// Tests no error
	inputNoErr := func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}

	Handler(inputNoErr).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Handler: %v, expected 200", w.Code)
	}

	Handler(nil).ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Handler: %v, expected 404", w.Code)
	}
}
