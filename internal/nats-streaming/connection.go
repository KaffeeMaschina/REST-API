package nats

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
)

func Connection(cluster, client, url string) {
	_, err := stan.Connect(cluster, client, stan.NatsURL(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to nats: %v\n", err)
		os.Exit(1)
	}

}
