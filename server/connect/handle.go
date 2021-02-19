package connect

import (
	"crypto/tls"
	"net"
)

func (rule *Rule) Handle(from net.Conn, initialBuffer []byte) {
	var (
		to  net.Conn
		err error
	)

	if rule.TLS.Enabled {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: rule.TLS.IgnoreBadCert,
		}
		to, err = tls.Dial(rule.Transport, rule.Address, tlsConfig)
	} else {
		to, err = net.Dial(rule.Transport, rule.Address)
	}

	if err != nil {
		return
	}

	to.Write(initialBuffer)

	go proxy(to, from)
	go proxy(from, to)
}

func proxy(from, to net.Conn) {
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
