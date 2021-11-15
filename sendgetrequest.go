package main

import (
	"io/ioutil"
	"net/http"
)

// Function that sends a GET request to a specified URL and returns the response
func sendGetRequest(targetAPIURL string) (response []byte, statusCode int, error error) {
	// Spawn the request
	req, _ := http.NewRequest("GET", targetAPIURL, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US")

	res, err := http.DefaultClient.Do(req)

	// If the request was not successfull log the error
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, res.StatusCode, nil

}
