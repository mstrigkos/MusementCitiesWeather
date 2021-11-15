package main

// global variable to store the configuration file parameters
var config map[string]interface{}

func main() {
	// Load app configuration from config.json
	loadConfig("config.json")
}
