package main

import (
	"flag"
	"github.com/vmware/vmware-go-kcl/logger"
)

// main entry point.
func main() {
	var endpoint string
	flag.StringVar(&endpoint, "endpoint", "localhost:8080", "endpoint for the backend workflow engine.")
	flag.Parse()

	config := logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logger.Info,
		ConsoleJSONFormat: false,
		//EnableFile:        true,
		//FileLevel:         logger.Info,
		//FileJSONFormat:    true,
		//Filename:          "log.log",
	}
	log := logger.NewLogrusLoggerWithConfig(config)

	log.Infof("Hello World")
}