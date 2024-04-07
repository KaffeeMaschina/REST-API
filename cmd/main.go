package main

import (
	"github.com/KaffeeMaschina/http-rest-api/internal/postgres"
)

func main() {
	postgres.Connectiondb()
	//nats.Connection("nats", "Nikita")
}
