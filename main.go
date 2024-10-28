// main.go
package main

import (
	"fmt"
	"github.com/miekg/dns"
	res "dns-server/resolve"
)

type dnsHandler struct{}

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		fmt.Printf("Received query: %s\n", question.Name)
		answers, err := res.Resolve(question.Name, question.Qtype)
		if err != nil {
			fmt.Printf("Error resolving domain: %v\n", err)
			continue
		}
		msg.Answer = append(msg.Answer, answers...)
	}

	w.WriteMsg(msg)
}

func main() {
	handler := new(dnsHandler)
	server := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: handler,
		UDPSize: 65535,
	}

	fmt.Println("Starting DNS server on port 53")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
