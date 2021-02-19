package listen

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/tomruk/multiplexor/server/connect"
	"github.com/tomruk/multiplexor/utils"
)

type Listener struct {
	Transport string `yml:"transport"`
	Address   string `yml:"address"`

	TLS struct {
		Enabled bool   `yml:"enabled"`
		Cert    string `yml:"cert"`
		Key     string `yml:"key"`
	} `yml:"tls"`

	rules []*connect.Rule
}

func (listener *Listener) Run(rules []*connect.Rule) error {
	var (
		lis net.Listener
		err error
	)

	listener.rules = rules

	if listener.TLS.Enabled {
		var tlsConfig *tls.Config
		tlsConfig, err = utils.GenerateTLSConfig(listener.TLS.Cert, listener.TLS.Key)
		if err != nil {
			return err
		}

		lis, err = tls.Listen(listener.Transport, listener.Address, tlsConfig)
	} else {
		lis, err = net.Listen(listener.Transport, listener.Address)
	}

	if err != nil {
		return err
	}

	fmt.Printf("Listening on: %s://%s", listener.Transport, listener.Address)

	if listener.TLS.Enabled {
		fmt.Println(" with TLS")
	} else {
		fmt.Println()
	}

	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				continue
			}

			go listener.handleConn(conn)
		}
	}()

	return nil
}
