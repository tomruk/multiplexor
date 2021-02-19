package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

var arguments struct {
	transport      *string
	listenAddress  *string
	connectAddress *string

	tls struct {
		enabled              *bool
		ignoreBadCertificate *bool
	}
}

var (
	listenTransport  string
	connectTransport string
)

func init() {
	arguments.transport = pflag.StringP("transport", "t", "tcp:tcp", "Transport protocols to use")
	arguments.listenAddress = pflag.StringP("listen", "l", "127.0.0.1:8080", "Bind to")
	arguments.connectAddress = pflag.StringP("connect", "c", "", "Connect to")

	arguments.tls.enabled = pflag.Bool("tls", true, "Use TLS for outgoing connection")
	arguments.tls.ignoreBadCertificate = pflag.Bool("ignore", false, "Ignore invalid certificate error")

	pflag.Parse()

	// Parse transport argument
	splitted := strings.Split(*arguments.transport, ":")
	if len(splitted) != 2 {
		fmt.Printf("Invalid transport: %s. must be ':' seperated\n", *arguments.transport)
	}
	listenTransport = splitted[0]
	connectTransport = splitted[1]

	// Check if --connect is empty
	if *arguments.connectAddress == "" {
		fmt.Println("Connect address cannot be empty. Please specify an address with -c/--connect option")
		os.Exit(1)
	}
}
