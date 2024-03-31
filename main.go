package main

import (
	"github.com/KaffeeMaschina/http-rest-api/pkg/client/postgresql"
)

func main() {
	postgresql.Init("postgres", "aaaa", "localhost", "5432", "testdb")
}
