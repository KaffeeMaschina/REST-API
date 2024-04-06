package main

import (
	"github.com/KaffeeMaschina/http-rest-api/internal/postgres"
)

func main() {
	postgres.Connection("postgres", "qwerty", "localhost", "5432", "postgres")
}
