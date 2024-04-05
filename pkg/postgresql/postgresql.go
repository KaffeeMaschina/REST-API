package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
)

type DB struct {
	conn *pgconn.PgConn
}

func Connection(username, password, host, port, database string) (*DB, error) {

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	c, err := pgconn.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err

	}

	return &DB{c}, nil
	//fmt.Printf("%v: connected to database\n", database)
}
