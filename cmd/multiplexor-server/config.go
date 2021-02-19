package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func init() {
	parseArgs()

	if *configPath != "" {
		viper.SetConfigFile(*configPath)
		viper.SetConfigType("yml")
	} else {
		executablePath, _ := os.Executable()
		executableDir := filepath.Dir(executablePath)

		viper.AddConfigPath(".")
		viper.AddConfigPath(executableDir)
		viper.AddConfigPath("/etc/multiplexor")
		viper.SetConfigName("server")
		viper.SetConfigType("yml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
