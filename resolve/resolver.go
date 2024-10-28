package resolve

import (
	"fmt"
	"github.com/miekg/dns"
)

func Resolve(domain string, qtype uint16) ([]dns.RR, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), qtype)
	m.RecursionDesired = true

	c := new(dns.Client)
	in, _, err := c.Exchange(m, "8.8.8.8:53") // Using Google's DNS server
	if err != nil {
		return nil, fmt.Errorf("DNS query failed: %v", err)
	}

	return in.Answer, nil
}
