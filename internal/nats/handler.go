package nats

import (
	"log"
	"os"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
	"github.com/nats-io/stan.go"
)

type StreamingHandler struct {
	conn  *stan.Conn
	sub   *Subscriber
	pub   *Publisher
	name  string
	isErr bool
}

func NewStreamingHandler(S *postgres.DB) *StreamingHandler {
	sh := StreamingHandler{}
	sh.Init(S)
	return &sh
}

func (sh *StreamingHandler) Init(S *postgres.DB) {
	sh.name = "StreamingHandler"
	err := sh.NewConnection()

	if err != nil {
		sh.isErr = true
		log.Printf("%s: StreamingHandler error: %s", sh.name, err)
	} else {
		sh.sub = NewSubscriber(S, sh.conn)
		sh.sub.Subscribe()

		sh.pub = NewPublisher(sh.conn)
		sh.pub.Publish()
	}
}
func (sh *StreamingHandler) NewConnection() error {
	sc, err := stan.Connect("test-cluster", "pub")
	if err != nil {
		os.Exit(2)

	}
	sh.conn = &sc

	return nil
}
