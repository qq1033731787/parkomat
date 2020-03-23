package dns

import (
	mdns "github.com/miekg/dns"
	"github.com/qq1033731787/parkomat/config"
	"strings"
)

type aHandler struct {
}

// Handle produces reply for A question
func (a *aHandler) Handle(msg *mdns.Msg, zone *config.Zone, question mdns.Question) (err error) {
	s := strings.Join(
		[]string{
			question.Name,
			"3600",
			"IN",
			"A",
			zone.A,
		}, " ")

	rr, err := mdns.NewRR(s)
	if err == nil {
		msg.Answer = append(msg.Answer, rr)
	}
	return
}
