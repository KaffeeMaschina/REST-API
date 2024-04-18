package main

import nats "github.com/KaffeeMaschina/http-rest-api/internal/nats/int"

func main() {
	nats.Subscriber()
	nats.Publisher()
}
