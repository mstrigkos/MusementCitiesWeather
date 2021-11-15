package main

import (
	"fmt"
	"log"
	"os"
)

// Function that can be called in order to save errors into a log file
func writeLog(msg string) {
	// Open the logfile as set in the configuration file
	file, err := os.OpenFile(config["logFile"].(string), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// If we cannot write the log file throw an error
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	log.SetOutput(file)
	log.Print(msg)

	file.Close()
}
