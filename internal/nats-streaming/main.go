package main

import "github.com/KaffeeMaschina/http-rest-api/internal/nats-streaming"

func main() {
	nats.Subscriber()
	nats.Publisher()
}
