package nats

import (
	"log"
	"os"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
	"github.com/nats-io/stan.go"
)

type StreamingHandler struct {
	conn      *stan.Conn
	sub       *Subscriber
	pub       *Publisher
	ClusterID string
	CLientID  string
	isErr     bool
}

func NewStreamingHandler(S *postgres.DB, ClusterID, CLientID string) *StreamingHandler {
	sh := StreamingHandler{
		ClusterID: ClusterID,
		CLientID:  CLientID,
	}
	sh.Init(S)
	return &sh
}

func (sh *StreamingHandler) Init(S *postgres.DB) {

	err := sh.NewConnection()

	if err != nil {
		sh.isErr = true
		log.Printf("StreamingHandler error: %s", err)
	} else {
		sh.sub = NewSubscriber(S, sh.conn)
		sh.sub.Subscribe()

		sh.pub = NewPublisher(sh.conn)
		sh.pub.Publish()
	}
}
func (sh *StreamingHandler) NewConnection() error {
	sc, err := stan.Connect(sh.ClusterID, sh.CLientID)
	if err != nil {
		os.Exit(2)

	}
	sh.conn = &sc

	return nil
}
