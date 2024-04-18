package nats

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func Publisher() {

	sc, err := stan.Connect("test-cluster", "pub")
	if err != nil {
		os.Exit(2)
	}
	orderData, err := json.Marschal(order)
	for i := 1; i <= 100; i++ {
		sc.Publish("Order", []byte("Order "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)

	}

}
