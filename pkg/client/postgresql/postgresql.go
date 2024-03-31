package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
)

func Init(username, password, host, port, database string) {

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	_, err := pgconn.Connect(context.Background(), dbUrl)
	if err != nil {

		fmt.Println("connection is closed")
	}
	fmt.Printf("%v: connected to database\n", database)
}
