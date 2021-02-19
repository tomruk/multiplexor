package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tomruk/multiplexor/server/connect"
	"github.com/tomruk/multiplexor/server/listen"

	"github.com/spf13/viper"
)

var (
	listeners []*listen.Listener
	rules     []*connect.Rule
	exitChan  = make(chan bool)
)

func main() {
	err := viper.UnmarshalKey("listen", &listeners)
	if err != nil {
		fmt.Printf("Cannot parse \"listen\" field of server.yml: %s\n", err.Error())
		os.Exit(1)
	}

	if len(listeners) == 0 {
		fmt.Println("No listener found in configuration, exiting")
		os.Exit(1)
	}

	err = viper.UnmarshalKey("rules", &rules)
	if err != nil {
		fmt.Printf("Cannot parse \"rules\" field of server.yml: %s\n", err.Error())
		os.Exit(1)
	}

	if len(rules) == 0 {
		fmt.Println("No rule found in configuration, exiting")
		os.Exit(1)
	}

	for _, rule := range rules {
		err = rule.Setup()
		if err != nil {
			fmt.Printf("Error at rule setup: %s\n", err.Error())
			os.Exit(1)
		}
	}

	for _, listener := range listeners {
		go runListener(listener)
	}

	<-exitChan
}

func runListener(listener *listen.Listener) {
	err := listener.Run(rules)
	if err != nil {
		fmt.Printf("Error while running listener: %s\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigc
		exitChan <- true
	}()
}
