package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Function that loads the config.json file
func loadConfig(configFile string) {
	// load configuration file with app parameters
	configJSON, err := ioutil.ReadFile(configFile)
	// If the configuration file does not exist or is not accessible terminate the app with error
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	parseConfig(configJSON)
}

// Function that parse configuration file contents into global variable config
func parseConfig(configJSON []byte) {
	// put app parameters in unstructured map
	jsonerr := json.Unmarshal(configJSON, &config)
	// If the configuration file is not a JSON format terminate the app with error
	if jsonerr != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", jsonerr)
		os.Exit(1)
	}

}
