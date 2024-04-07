package config

import (
	"os"
)

func Config() {
	os.Setenv("USERNAME_DB", "postgres")
	os.Setenv("PASSWORD_DB", "qwerty")
	os.Setenv("HOST_DB", "localhost:5432")
	os.Setenv("DBNAME_DB", "postgres")

}
