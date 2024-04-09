package nats

import "github.com/nats-io/stan.go"

type Subscriber struct {
	conn *stan.Conn
}

func NewSub() {
	//Sub := Subscriber{conn: }
}

//func SubscriptionNS() {
//	sub, err := (client).Subscribe
//}
