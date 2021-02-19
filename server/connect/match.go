package connect

import (
	"fmt"

	"github.com/tomruk/multiplexor/server/match"
)

// Sets the MatchFunc with the appropriate match function
// Returns error if match name is invalid
func (rule *Rule) setMatch() error {
	switch rule.Match {
	default:
		return fmt.Errorf("invalid match: %s", rule.Match)
	case "http":
		rule.MatchFunc = match.MatchHTTP
	case "socks4":
		rule.MatchFunc = match.MatchSOCKS4
	case "socks5":
		rule.MatchFunc = match.MatchSOCKS5
	case "ssh":
		rule.MatchFunc = match.MatchSSH
	case "openvpn":
		rule.MatchFunc = match.MatchOpenVPN
	}
	return nil
}
