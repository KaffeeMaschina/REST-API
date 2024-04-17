package main

import (
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

	for i := 1; ; i++ {
		sc.Publish("Order", []byte("Order "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}
}
