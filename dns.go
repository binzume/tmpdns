package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/miekg/dns"
)

func rr(name, qtype, value string) dns.RR {
	r, _ := dns.NewRR(fmt.Sprintf("%s 1 %s %s", name, strings.ToUpper(qtype), value))
	return r
}

func main() {
	port := flag.String("p", "53", "bind to [addr:]port")
	zone := flag.String("z", ".", "zone")
	flag.Parse()

	records := map[string][]dns.RR{}
	for _, v := range flag.Args() {
		r := strings.SplitN(v, ":", 3)
		if len(r) == 3 { // name:type:value
			if !strings.HasSuffix(r[0], ".") {
				r[0] = r[0] + "." + *zone
			}
			records[r[0]] = append(records[r[0]], rr(r[0], r[1], r[2]))
		} else {
			log.Printf("unknown record %s", v)
		}
	}

	dns.HandleFunc(*zone, func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		switch r.Opcode {
		case dns.OpcodeQuery:
			for _, q := range m.Question {
				log.Printf("query %s %d\n", q.Name, q.Qtype)
				for _, rr := range records[strings.ToLower(q.Name)] {
					if rr.Header().Rrtype == q.Qtype {
						log.Printf("RR %v\n", rr)
						m.Answer = append(m.Answer, rr)
					}
				}
			}
		}
		w.WriteMsg(m)
	})

	if ! strings.Contains(*port, ":") {
		*port = ":" + *port
	}

	server := &dns.Server{Addr: *port, Net: "udp"}
	log.Printf("Start DNS Server at %s  %v\n", *port, records)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
