package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Test GetCities function
func TestGetCities(t *testing.T) {
	// Format dummy data to the expected type of unmarshaled response
	expected := []City{{"Amsterdam", 12.55, 15.22}, {"London", 23.19, -42.05}}

	// Initiate a mock server to respond to the request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// encode the expected data
		response, _ := json.Marshal(expected)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		w.Write(response)
	}))

	defer srv.Close()

	// set the mock server url on the global config variable
	configJson := `{"musementAPIURL": "` + srv.URL + `"}`
	json.Unmarshal([]byte(configJson), &config)
	cities := []City{}
	getCities(&cities)

	// We expect the unmarshaled json to be equal with the initial data
	if !(reflect.DeepEqual(cities, expected)) {
		t.Errorf("expected cities slice of objects to be equal with the unmarshaled server json response")
	}
}
