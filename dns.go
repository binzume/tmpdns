package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/miekg/dns"
)

var records = map[string]dns.RR{}

func handleQuery(m *dns.Msg) {
	for _, q := range m.Question {
		log.Printf("query %s %s\n", q.Name, q.Qtype)
		name := strings.ToLower(q.Name)
		qtype := ""
		switch q.Qtype {
		case dns.TypeA:
			qtype = "a"
		case dns.TypeAAAA:
			qtype = "aaaa"
		case dns.TypeTXT:
			qtype = "txt"
		case dns.TypeCNAME:
			qtype = "cname"
		case dns.TypeNS:
			qtype = "ns"
		}
		if rr, ok := records[name+":"+qtype]; ok {
			log.Printf("RR %v\n", rr)
			m.Answer = append(m.Answer, rr)
		}
	}
}

func rr(name, qtype, value string) dns.RR {
	r, _ := dns.NewRR(fmt.Sprintf("%s 1 %s %s", name, strings.ToUpper(qtype), value))
	return r
}

func main() {
	port := flag.Int("p", 53, "listen port")
	zone := flag.String("z", ".", "zone")
	flag.Parse()

	for _, v := range flag.Args() {
		r := strings.SplitN(v, ":", 3)
		if len(r) == 3 { // name:type:value
			if !strings.HasSuffix(r[0], ".") {
				r[0] = r[0] + "." + *zone
			}
			records[r[0]+":"+r[1]] = rr(r[0], r[1], r[2])
		} else {
			log.Printf("unknown record %s", v)
		}
	}

	dns.HandleFunc(*zone, func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		switch r.Opcode {
		case dns.OpcodeQuery:
			handleQuery(m)
		}
		w.WriteMsg(m)
	})

	server := &dns.Server{Addr: ":" + strconv.Itoa(*port), Net: "udp"}
	log.Printf("Start DNS Server at :%d  %v\n", *port, records)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
