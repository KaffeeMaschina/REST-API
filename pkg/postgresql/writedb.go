package postgresql

import (
	"context"

	"github.com/jackc/pgconn"
)

func Write(db *DB, sql string) {
	a := pgconn.Exec(context.Background(), sql)
	return a
}
