package nats

import (
	"encoding/json"
	"log"
	"os"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage"
	"github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	sub stan.Subscription
	BD  *postgres.DB
	sc  *stan.Conn
}

func NewSubscriber(db *postgres.DB, conn *stan.Conn) *Subscriber {
	return &Subscriber{
		BD: db,
		sc: conn,
	}
}

func (s *Subscriber) Subscribe() {
	var err error
	s.sub, err = (*s.sc).Subscribe("Order", func(m *stan.Msg) {
		//fmt.Printf("Got: %s\n", string(m.Data))
		s.MsgToStorage(m.Data)
	},
		stan.DurableName("my-durable"),
		stan.DeliverAllAvailable())

	if err != nil {
		os.Exit(11)
	}

}

func (s *Subscriber) MsgToStorage(data []byte) {
	ReceivedMsg := storage.Orders{}
	err := json.Unmarshal(data, &ReceivedMsg)
	if err != nil {
		log.Printf("Unmarschal error: %s", err)
	}
	//fmt.Println(ReceivedMsg)
	s.BD.AddOrder(ReceivedMsg)
}
func (s *Subscriber) Unsubscribe() {
	if s.sub != nil {
		s.sub.Unsubscribe()
	}
}
