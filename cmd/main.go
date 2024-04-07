package main

import (
	"os"

	"github.com/KaffeeMaschina/http-rest-api/config"
	"github.com/KaffeeMaschina/http-rest-api/internal/postgres"
)

func main() {
	config.Config()
	postgres.Connectiondb(os.Getenv("USERMANE_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("HOST_DB"), os.Getenv("PORT_DB"), os.Getenv("DBNAME_DB"))
	//nats.Connection("nats", "Nikita")
}
