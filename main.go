package main

type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// global variable to store the configuration file parameters
var config map[string]interface{}

func main() {
	// Load app configuration from config.json
	loadConfig("config.json")
}
