package main

import (
	"github.com/KaffeeMaschina/http-rest-api/internal/nats-streaming"
	"github.com/KaffeeMaschina/http-rest-api/internal/postgres"
)

func main() {
	postgres.Connectiondb("postgres", "qwerty", "localhost", "5432", "postgres")
	nats.Connection("test-cluster", "NC27N2A4LCNWMGZGKLG2S2WRE527W4SPEEBQMTZGVWAZWZQCD6SPNVOC", "4222")
}
