package match

import (
	"regexp"
)

var openVPNRegexes = []*regexp.Regexp{
	regexp.MustCompile(`^\x00[\x0d-\xff]\x38`),
	regexp.MustCompile(`^\x00[\x0d-\xff]$`),
}

func MatchOpenVPN(buffer []byte) bool {
	for _, re := range openVPNRegexes {
		if re.Match(buffer) {
			return true
		}
	}
	return false
}
