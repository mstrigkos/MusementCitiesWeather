package main

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WeatherForecast struct {
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// global variable to store the configuration file parameters
var config map[string]interface{}

func main() {
	// Load app configuration from config.json
	loadConfig("config.json")

	cities := []City{}

	// Fetch the cities from the Musement API
	getCities(&cities)
}
