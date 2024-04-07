package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connectiondb(username, password, host, port, database string) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	dbconn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close()

	var sql string
	err = dbconn.QueryRow(context.Background(), "select name from delivery").Scan(&sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(2)
	}

	fmt.Println(sql)
}
