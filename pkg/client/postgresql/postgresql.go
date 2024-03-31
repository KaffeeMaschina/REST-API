package postgresql

import (
	"context"

	"github.com/jackc/pgconn"
)

type client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
}
