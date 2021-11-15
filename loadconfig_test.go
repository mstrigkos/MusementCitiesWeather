package main

import (
	"fmt"
	"testing"
)

// Test ParseConfig
func TestParseConfig(t *testing.T) {
	// Expected map
	expected := map[string]string{"musementAPIURL": "https://musement.test", "logFile": "errors.log", "weatherAPIURL": "https://weatherapi.test", "weatherAPIKey": "12345"}
	// JSON data mock of the configuration file
	configFileContents := `{"musementAPIURL": "https://musement.test","logFile": "errors.log","weatherAPIURL": "https://weatherapi.test","weatherAPIKey": "12345"}`
	parseConfig(([]byte(configFileContents)))
	// we use fmt.Sprint instead of reftect becaulse config is a map interface
	if fmt.Sprint(expected) != fmt.Sprint(config) {
		t.Errorf("configuration map was different that the one expected")
	}

}
