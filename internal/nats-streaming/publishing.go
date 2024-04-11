package nats

import (
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func Publish() {
	sc, _ := stan.Connect("test-cluster", "Pub")
	defer sc.Close()
	err := sc.Publish("foo", []byte("Hello World"))
	_ = err

	for i := 1; ; i++ {
		sc.Publish("Order", []byte("Order "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}
}
