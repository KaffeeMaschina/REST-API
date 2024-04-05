package main

import (
	"github.com/KaffeeMaschina/http-rest-api/pkg/postgres"
)

func main() {
	postgres.Connection("postgres", "qwerty", "localhost", "5432", "postgres")

}
