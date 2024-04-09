package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(username, password, host, port, database string) (*Storage, error) {
	const op = "storage.postgres.New"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var sql string = "SELECT NAME FROM DELIVERY"
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(2)
	}
	defer rows.Close()

	fmt.Println(sql)
	return &Storage{db: db}, nil
}

//func Connectiondb(username, password, host, port, database string) {

//}
