package listen

import (
	"net"
	"time"

	"github.com/tomruk/multiplexor/server/connect"
)

func (listener *Listener) handleConn(conn net.Conn) {
	initialBuffer := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	n, err := conn.Read(initialBuffer)
	if err != nil {
		conn.Close()
		return
	}

	conn.SetReadDeadline(time.Time{})
	initialBuffer = initialBuffer[:n]

	var rule *connect.Rule
	for _, r := range listener.rules {
		if r.MatchFunc(initialBuffer) {
			rule = r
			break
		}
	}

	if rule != nil {
		rule.Handle(conn, initialBuffer)
		return
	}
	conn.Close()
}
