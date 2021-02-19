package match

import (
	"bytes"
	"net/http"
	"strings"
)

var defaultHTTPMethods = [][]byte{
	[]byte("POST"),
	[]byte("GET"),
	[]byte("OPTIONS"),
	[]byte("HEAD"),
	[]byte("PUT"),
	[]byte("DELETE"),
	[]byte("TRACE"),
	[]byte("CONNECT"),
}

func MatchHTTP(buffer []byte) bool {
	if !findAnHTTPMethod(buffer) {
		return false
	}

	_, proto, ok := parseRequestLine(string(buffer))
	if !ok {
		return false
	}

	version, _, ok := http.ParseHTTPVersion(proto)
	return ok && version == 1
}

func findAnHTTPMethod(buffer []byte) bool {
	for i := 0; i < len(defaultHTTPMethods); i++ {
		if bytes.HasPrefix(buffer, defaultHTTPMethods[i]) {
			return true
		}
	}
	return false
}

// From net/http
func parseRequestLine(line string) (method, proto string, ok bool) {
	s1 := strings.Index(line, " ")
	s2 := strings.Index(line[s1+1:], " ")
	s3 := strings.Index(line, "\n")

	if s1 < 0 || s2 < 0 {
		return
	}
	s2 += s1 + 1
	if s3 <= s2+1 {
		return
	}

	return line[:s1], line[s2+1 : s3-1], true
}
