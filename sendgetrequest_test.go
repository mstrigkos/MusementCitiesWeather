package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// Test SendGetRequest function
func TestSendGetRequest(t *testing.T) {
	// String expected value
	expected := "response ok"
	// Initiate a mock server to respond to the request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// encode the mocked data
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		w.Write([]byte(expected))
	}))

	defer srv.Close()

	res, statusCode, err := sendGetRequest(srv.URL)

	// Check for request error or invalid response code
	if err != nil {
		t.Errorf("http request returned an error")
	} else if statusCode != 200 {
		t.Errorf("server did not return an 200 OK response code")
	} else if !(reflect.DeepEqual(strings.TrimSpace(expected), strings.TrimSpace(string(res)))) {
		t.Errorf("channel output was not as expected")
	}

}
