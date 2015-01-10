package main

import (
	"encoding/json"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	input := &Route{
		path:    "/testing",
		handler: nil,
	}

	b, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Route Marshal: %v", err)
	}

	if string(b) != `"/testing"` {
		t.Errorf(`Route Marshal: %s, expected "/testing"`, b)
	}
}
