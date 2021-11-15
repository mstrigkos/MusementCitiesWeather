package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Function that gets the forecast from Weather API for a specific city and sends a message to the channel
func getWeather(city City, ch chan<- string) {
	// Format the weather API url adding the API key
	url := config["weatherAPIURL"].(string) + "?key=" + config["weatherAPIKey"].(string)

	// Transform and add the lat/long parameters to the url
	url += "&q=" + strings.Replace(fmt.Sprintf("%f", city.Latitude), ".", ",", 1) + "/" + strings.Replace(fmt.Sprintf("%f", city.Longitude), ".", ",", 1)
	url += "&days=2&aqi=no&alerts=no"
	forecasts := WeatherForecast{}
	output := ""
	// Send request to weather API
	res, statusCode, err := sendGetRequest(url)
	if err != nil {
		writeLog(err.Error())
	} else if statusCode != 200 {
		writeLog("Weather API error status code " + fmt.Sprintf("%d", statusCode))
	} else {
		json.Unmarshal(res, &forecasts)
		// Format the output that will be sent to STDOUT
		output = "Processed city " + city.Name + " | "
		for i, forecast := range forecasts.Forecast.Forecastday {
			output += forecast.Day.Condition.Text
			if i == 0 {
				output += " - "
			}
		}
	}

	// Send the output value to the channel
	ch <- fmt.Sprintf("%s\n", output)
}
