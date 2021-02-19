package match

import (
	"bytes"
)

var sshPrefix = []byte("SSH-")

func MatchSSH(buffer []byte) bool {
	return bytes.HasPrefix(buffer, sshPrefix)
}
