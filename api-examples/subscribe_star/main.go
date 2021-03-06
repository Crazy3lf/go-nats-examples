package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin subscribe_star]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for 2 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Subscribe
	if _, err := nc.Subscribe("time.*.east", func(m *nats.Msg) {
		log.Printf("%s: %s", m.Subject, m.Data)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for the 2 messages to come in
	wg.Wait()

	// [end subscribe_star]
}
