package nats

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
)

type nStreaming struct {
	conn *stan.Conn
	sub  *Subscriber
	//pub *Publisher
}

var ns *nStreaming

func (ns *nStreaming) ConnectionNS(clusterID, clientID string) (client *stan.Conn, err error) {

	сlient, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Nats connection failed: %v\n", err)
		os.Exit(2)
		return nil, err
	} else {
		fmt.Println("Connected to nats")
	}
	return &сlient, nil

}
