package nats

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
)

func Subscriber() {

	sc, errConn := stan.Connect("test-cluster", "sub")
	if errConn != nil {
		os.Exit(10)
	}
	defer sc.Close()
	sub, errSub := sc.Subscribe("Order", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
	},

		stan.DeliverAllAvailable())
	sub.IsValid()
	if errSub != nil {
		os.Exit(11)
	}

}
