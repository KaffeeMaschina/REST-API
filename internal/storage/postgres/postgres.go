package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	mdbUrl := dbUrl + "?sslmode=disable"
	migrationPath := "file:///Users/nikitacode/LearnGo/Service/rest-api/schema"
	fmt.Printf("migrationPath : %s\n", migrationPath)
	m, err := migrate.New(
		migrationPath,
		mdbUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New")

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to apply")

		}
		panic(err)
	}
	fmt.Println("migrations applied successfully")
	return &Storage{db: db}, nil
}

/*//var sql string =
	rows, err := db.Query(context.Background(), `CREATE TABLE IF NOT EXISTS PAYMENT (
	PAYMENT_ID INT PRIMARY KEY NOT NULL,
	"TRANSACTION"  VARCHAR(128),
	REQUEST_ID VARCHAR(128),
	CURRENCY VARCHAR(128),
	"PROVIDER" VARCHAR(128),
	AMOUNT INT,
	PAYMENT_DT INT,
	BANK VARCHAR(128),
	DELIVERY_COST INT,
	GOODS_TOTAL INT,
	CUSTOM_FEE INT
	)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(2)
	}
	defer rows.Close()

	return &Storage{db: db}, nil
}

//func Connectiondb(username, password, host, port, database string) {

//}*/
