package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connectiondb() {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s", os.Getenv("USERNAME_DB"), os.Getenv("PASSWORD_DB"), os.Getenv("HOST_DB"), os.Getenv("DBNAME_DB"))

	dbconn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbconn.Close()

	var sql string
	err = dbconn.QueryRow(context.Background(), "select delivery_id from delivery").Scan(&sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(sql)
}
