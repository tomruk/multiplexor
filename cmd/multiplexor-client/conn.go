package main

import (
	"crypto/tls"
	"net"
)

func connect(transport, addr string, useTLS, ignoreBadCert bool) (conn net.Conn, err error) {
	if useTLS {
		conn, err = tls.Dial(transport, addr, &tls.Config{InsecureSkipVerify: ignoreBadCert})
	} else {
		conn, err = net.Dial(transport, addr)
	}
	return
}

func proxy(from net.Conn, to net.Conn) {
	buffer := make([]byte, 4096)

	for {
		n, err := from.Read(buffer)
		if err != nil {
			return
		}

		_, err = to.Write(buffer[:n])
		if err != nil {
			return
		}
	}
}
