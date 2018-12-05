package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/go-nats"
)

// Constants
const (
	defaultURL = "nats://localhost:4222"
)

var debug bool

func printf(format string, v ...interface{}) {
	out := fmt.Sprintf(format, v...)
	log.Printf("%v: %s", time.Now().Format("2016-04-08 15:04:05.00"), out)
}

func debugf(format string, v ...interface{}) {
	if debug {
		printf(format, v...)
	}
}

func main() {
	var (
		URL     string
		subj    string
		user    string
		pass    string
		jwt     string
		nk      string
		chain   string
		nc      *nats.Conn
		err     error
		start   time.Time
		elapsed time.Duration
	)

	flag.StringVar(&URL, "s", defaultURL, "default url (nats://localhost:4222)")
	flag.StringVar(&subj, "subj", "", "Publish subject, default is unset")
	flag.StringVar(&user, "user", "", "Subscribe subject, default is unset")
	flag.StringVar(&pass, "pass", "", "name of client, default is hostname-pid")
	flag.StringVar(&jwt, "jwt", "", "User JWT path (requires nk), default is unset")
	flag.StringVar(&nk, "nk", "", "User nkey path (requires jwt), default is unset")
	flag.StringVar(&chain, "chain", "", "User chained file (JWT and nkey), default is unset")
	flag.BoolVar(&debug, "debug", false, "enable debugging")

	log.SetFlags(0)
	flag.Parse()

	start = time.Now()
	if chain != "" {
		debugf("Using chain file for authentication.\n")
		nc, err = nats.Connect(URL,
			nats.UserCredentials(chain))
	} else if jwt != "" {
		debugf("Using jwt/nk files for authentication.\n")
		nc, err = nats.Connect(URL,
			nats.UserCredentials(jwt, nk))
	} else if user != "" {
		debugf("Using user/pass for authentication.\n")
		nc, err = nats.Connect(URL,
			nats.UserInfo(user, pass))
	} else {
		debugf("No connection authentication.\n")
		nc, err = nats.Connect(URL)
	}
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	printf("Connect time: %v.", time.Since(start))

	if subj == "" {
		subj = nats.NewInbox()
	}

	var recvWg sync.WaitGroup
	recvWg.Add(1)

	_, err = nc.Subscribe(subj, func(msg *nats.Msg) {
		elapsed = time.Since(start)
		recvWg.Done()
	})
	if err != nil {
		log.Fatalf("couldn't subscribe: %v", err)
	}

	start = time.Now()
	if err := nc.Publish(subj, []byte("testping")); err != nil {
		log.Fatalf("couldn't publish: %v", err)
	}
	if err := nc.Flush(); err != nil {
		log.Fatalf("couldn't flush: %v", err)
	}

	recvWg.Wait()

	printf("Ping time:    %v.", elapsed.String())
}
