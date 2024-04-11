package nats

import (
	"fmt"
	"sync"

	"github.com/nats-io/stan.go"
)

/*type nStreaming struct {
	conn *stan.Conn
	//sub  *Subscriber
	//pub *Publisher
}*/

//var ns *nStreaming

func /*(ns *nStreaming)*/ ConnectionNS( /*clusterID, clientID string) (client *stan.Conn, err error*/ ) {
	sc, _ := stan.Connect("test-cluster", "Nikita")
	defer sc.Close()

	sc.Subscribe("Order", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
	},
		stan.DeliverAllAvailable())
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()

	sc.Subscribe("foo",
		func(m *stan.Msg) {
			fmt.Printf("Got: %s\n", string(m.Data))
		},
		stan.DeliverAllAvailable())

	/*сlient, err := stan.Connect(
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

	func NewPub(client stan.Conn) error {
		err := client.Publish("foo", []byte("Hello World"))
		if err != nil {
			return err
		}
	*/
}
