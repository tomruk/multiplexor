package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen(listenTransport, *arguments.listenAddress)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Listening on: %s://%s\n", listenTransport, *arguments.listenAddress)
	fmt.Printf("Connections will be proxied to: %s://%s\n", connectTransport, *arguments.connectAddress)

	for {
		from, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		go handleConn(from)
	}
}

func handleConn(from net.Conn) {
	to, err := connect(
		connectTransport,
		*arguments.connectAddress,
		*arguments.tls.enabled,
		*arguments.tls.ignoreBadCertificate,
	)

	if err != nil {
		from.Close()
		return
	}

	go proxy(from, to)
	go proxy(to, from)
}
