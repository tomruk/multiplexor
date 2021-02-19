package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	configPath   *string
	generateCert *bool
)

func parseArgs() {
	configPath = pflag.StringP("config", "c", "", "Configuration file")
	generateCert = pflag.BoolP("generate", "g", false, "Generate a self signed SSL certificate")
	pflag.Parse()

	if *generateCert == true {
		err := generateSelfSignedTLSCert()
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
}
