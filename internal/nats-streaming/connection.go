package nats

import (
	"github.com/nats-io/stan.go"
)

func Connection(cluster, client string) {
	conn := stan.Connect(cluster)

}
