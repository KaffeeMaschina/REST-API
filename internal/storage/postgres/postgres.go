package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage"
	"github.com/golang-migrate/migrate/v4"

	// postgres driver for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// file driver for migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	db  *pgxpool.Pool
	csh *Cache
}

// Create a database
func New(username, password, host, port, database string, csh *Cache) (*DB, error) {

	const op = "storage.postgres.New"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	S := &DB{db: db,
		csh: csh,
	}

	Migration(dbUrl)
	S.SetCacheFromDB()

	return S, nil
}

// Read and run migrations
func Migration(dbUrl string) {
	mdbUrl := dbUrl + "?sslmode=disable"
	migrationPath := "file://db/migrations"

	m, err := migrate.New(
		migrationPath,
		mdbUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		fmt.Println("no migration to apply")

	} else {
		fmt.Println("migrations applied successfully")
	}
}

// Write data to database from Nats
func (S *DB) AddOrderByOID(o storage.Orders) {

	oid := S.GetOID()
	//Check if oid is absent, then write to database
	if oid == "" {
		S.AddOrder(o)
	}
}

// Write data to database
func (S *DB) AddOrder(o storage.Orders) error {

	var lastInsertId int64
	var lastInsertOrderUid string

	tx, err := S.db.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer tx.Rollback(context.Background())
	//Write items to database
	for _, item := range o.Items {
		err = tx.QueryRow(context.Background(), `INSERT INTO items (chrt_id, track_number, price, rid, name_, sale, size_, 
		total_price, nm_id, brand, status_) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING chrt_id`,
			item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale, item.Size,
			item.Total_price, item.Nm_id, item.Brand, item.Status).Scan(&lastInsertId)
	}
	if err != nil {
		log.Printf("unable to insert data (items): %v\n", err)

	}
	//Write deliveries to databse
	ItemIdfk := lastInsertId
	err = tx.QueryRow(context.Background(), `INSERT INTO deliveries (delivery_id, name_, phone, zip, city, address_, region, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING delivery_id`,
		o.Delivery.Delivery_id, o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip, o.Delivery.City, o.Delivery.Address,
		o.Delivery.Region, o.Delivery.Email).Scan(&lastInsertId)
	if err != nil {
		log.Printf("unable to insert data (deliveries): %v\n", err)

	}
	DeliveryIdfk := lastInsertId
	//Write payments to database
	err = tx.QueryRow(context.Background(), `INSERT INTO payments (payment_id, transactions, request_id, currency, provider_, 
		amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING payment_id`,
		o.Payments.Payment_id, o.Payments.Transaction, o.Payments.Request_id, o.Payments.Currency, o.Payments.Provider,
		o.Payments.Amount, o.Payments.Payment_dt, o.Payments.Bank, o.Payments.Delivery_cost, o.Payments.Goods_total, o.Payments.Custom_fee).Scan(&lastInsertId)
	if err != nil {
		log.Printf("unable to insert data (payments): %v\n", err)

	}
	PaymentId_fk := lastInsertId
	//Write orders to database
	err = tx.QueryRow(context.Background(), `INSERT INTO orders (order_uid, track_number, name_entry, locale, internal_signature,
		customer_id, delivery_service, shardkey, sm_id, date_created, off_shard, fk_delivery_id, fk_payment_id, fk_item_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING order_uid`,
		o.Order_uid, o.Track_number, o.Entry, o.Locale, o.Internal_signature, o.Customer_id, o.Delivery_service, o.Shardkey,
		o.Sm_id, o.Date_created, o.Oof_shard, DeliveryIdfk, PaymentId_fk, ItemIdfk).Scan(&lastInsertOrderUid)
	if err != nil {
		log.Printf("unable to insert data (orders): %v\n", err)

	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	S.SetCacheFromDB()

	return nil
}

// Write data to cache from database
func (S *DB) SetCacheFromDB() {

	oid := S.GetOID()

	o, err := S.GetOrderFromDB(oid)
	if err != nil {
		log.Printf(": unable to get orders from database: %v\n", err)
	}

	S.csh.SetCache(oid, o)

}

// Select data from database to storage.Orders{}
func (S *DB) GetOrderFromDB(oid string) (o storage.Orders, err error) {

	var payment_id_fk, delivery_id_fk, item_id_fk int64

	//Select data from orders
	err = S.db.QueryRow(context.Background(), `SELECT order_uid, track_number, name_entry, locale, internal_signature, customer_id, 
	delivery_service, shardkey, sm_id, date_created, off_shard, fk_delivery_id, fk_payment_id, fk_item_id 
	FROM orders WHERE order_uid = $1`,
		oid).Scan(&o.Order_uid, &o.Track_number, &o.Entry, &o.Locale, &o.Internal_signature, &o.Customer_id, &o.Delivery_service, &o.Shardkey,
		&o.Sm_id, &o.Date_created, &o.Oof_shard, &payment_id_fk, &delivery_id_fk, &item_id_fk)
	if err != nil {
		log.Printf(": unable to get orders from database: %v\n", err)
		return o, err
	}
	//Select data from payments
	err = S.db.QueryRow(context.Background(), `SELECT payment_id, transactions, request_id, currency, provider_, amount, payment_dt,
	bank, delivery_cost, goods_total, custom_fee FROM payments WHERE payment_id = $1`,
		payment_id_fk).Scan(&o.Payments.Payment_id, &o.Payments.Transaction, &o.Payments.Request_id, &o.Payments.Currency,
		&o.Payments.Provider, &o.Payments.Amount, &o.Payments.Payment_dt, &o.Payments.Bank, &o.Payments.Delivery_cost,
		&o.Payments.Goods_total, &o.Payments.Custom_fee)
	if err != nil {
		log.Printf(": unable to get payment from database: %v\n", err)
		return o, err
	}
	//Select data from deliveries
	err = S.db.QueryRow(context.Background(), `SELECT delivery_id, name_, phone, zip, city, address_, region, email
	FROM deliveries WHERE delivery_id = $1`,
		delivery_id_fk).Scan(&o.Delivery.Delivery_id, &o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, &o.Delivery.City,
		&o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email)
	if err != nil {
		log.Printf(": unable to get deliveries from database: %v\n", err)
		return o, err
	}
	//Select data from items
	rows, err := S.db.Query(context.Background(), `SELECT * FROM items WHERE chrt_id = $1`, item_id_fk)
	if err != nil {
		return o, err
	}
	defer rows.Close()

	for rows.Next() {
		var item storage.Items
		fmt.Println(item)
		if err := rows.Scan(&item.Chrt_id, &item.Track_number, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size,
			&item.Total_price, &item.Nm_id, &item.Brand, &item.Status); err != nil {
			return o, err
		}

		o.Items = append(o.Items, item)

		if err = rows.Err(); err != nil {
			return o, err
		}
	}
	return o, nil
}

// Get data for server
func (S *DB) OrderOut() (o storage.Orders) {

	oid := S.GetOID()

	o = S.csh.OrderOutCache(oid)

	return
}

// Select order_uid from database
func (S *DB) GetOID() (oid string) {

	err := S.db.QueryRow(context.Background(), `SELECT order_uid FROM orders`).Scan(&oid)
	if err != nil {
		log.Printf(": unable to get order_id from database: %v\n", err)
	}
	return oid
}
