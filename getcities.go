package main

import (
	"encoding/json"
	"fmt"
)

// Function to get the available cities from Musement API
func getCities(cities *[]City) {
	// Send request to Musement API
	res, statusCode, err := sendGetRequest(config["musementAPIURL"].(string))

	// Check for error and log if any
	if err != nil {
		writeLog(err.Error())
		// Check if the HTTP status code is not OK and log the error code
	} else if statusCode != 200 {
		writeLog("Musiment API error status code " + fmt.Sprintf("%d", statusCode))
	} else {
		json.Unmarshal(res, cities)
	}
}
