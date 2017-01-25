package main

import (
	"fmt"
	"log"

	"os"

	"github.com/hashicorp/serf/serf"
)

var (
	nodes = []string{
		"10.20.30.41", "10.20.30.42", "10.20.30.43",
	}

	serfReconcile = make(chan serf.Member, 32)
	serfEvents    = make(chan serf.Event, 128)
)

func main() {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.BindAddr = os.Args[1]
	conf.MemberlistConfig.BindPort = 12345
	conf.MemberlistConfig.AdvertiseAddr = conf.MemberlistConfig.BindAddr
	conf.MemberlistConfig.AdvertisePort = conf.MemberlistConfig.BindPort
	conf.EventCh = serfEvents
	conf.Tags["role"] = "goat"
	fmt.Printf("%#v\n", conf)

	s, err := serf.Create(conf)
	fatalError(err)

	fmt.Printf("%#v\n", s.Memberlist())

	if len(os.Args) > 2 {
		nNodes, err := s.Join(os.Args[2:], true)
		fatalError(err)
		log.Println("Nodes:", nNodes)
	}

	for ev := range serfEvents {
		switch ev.EventType() {
		case serf.EventMemberJoin:
			fmt.Printf("Member Join: %#v\n", ev)
		case serf.EventMemberFailed:
			fmt.Printf("Member Failed: %#v\n", ev)
		case serf.EventMemberLeave:
			fmt.Printf("Member Leave: %#v\n", ev)
		case serf.EventMemberReap:
			fmt.Printf("Member Reap: %#v\n", ev)
		case serf.EventMemberUpdate:
			fmt.Printf("Member Update: %#v\n", ev)
		}
	}
}

func fatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
