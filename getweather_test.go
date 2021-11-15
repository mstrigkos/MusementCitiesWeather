package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// Test GetCities function
func TestGetWeather(t *testing.T) {
	// Format weather API dummy data to the expected type of unmarshaled response
	expected := "Processed city Amsterdam | Partly cloudy - Sunny"
	weatherAPIJSONmock := `{"forecast":{"forecastday":[{"date": "2021-11-11","day": {"condition": {"text": "Partly cloudy"}}},{"date": "2021-11-12","day": {"condition": {"text": "Sunny"}}}]}}`
	// Initiate a mock server to respond to the request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// encode the mocked data
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		w.Write([]byte(weatherAPIJSONmock))
	}))

	defer srv.Close()

	// set the mock server url on the global config variable
	configJson := `{"weatherAPIURL": "` + srv.URL + `", "weatherAPIKey": "dummyKey"}`
	json.Unmarshal([]byte(configJson), &config)
	city := City{"Amsterdam", 12.55, 15.22}
	ch := make(chan string)
	go getWeather(city, ch)

	channelResponse := <-ch
	// We expect the unmarshaled json to be equal with the initial data
	if !(reflect.DeepEqual(strings.TrimSpace(channelResponse), strings.TrimSpace(expected))) {
		t.Errorf("channel output was not as expected")
	}
}
