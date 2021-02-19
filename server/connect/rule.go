package connect

import (
	"github.com/tomruk/multiplexor/server/match"
)

type Rule struct {
	Transport string `yml:"transport"`
	Address   string `yml:"address"`

	TLS struct {
		Enabled       bool `yml:"enabled"`
		IgnoreBadCert bool `yml:"ignoreBadCert"`
	} `yml:"tls"`

	Match     string `yml:"match"`
	MatchFunc match.MatchFunc
}

func (rule *Rule) Setup() error {
	return rule.setMatch()
}
