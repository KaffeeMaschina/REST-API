package main

import (
	"github.com/KaffeeMaschina/http-rest-api/pkg/postgresql"
)

func main() {
	postgresql.Connection("postgres", "qwerty", "", "5432", "testdb")

}
