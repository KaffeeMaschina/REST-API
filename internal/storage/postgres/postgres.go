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
func (S *DB) AddOrder(o storage.Orders) (string, error) {
	const op = "storage.postgres.AddOrder"
	var lastInsertId int64

	var lastInsertOrderUid string

	tx, err := S.db.Begin(context.Background())
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(context.Background())
	for _, item := range o.Items {
		err = tx.QueryRow(context.Background(), `INSERT INTO items (chrt_id, track_number, price, rid, name_, sale, size_, 
		total_price, nm_id, brand, status_) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING chrt_id`,
			item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale, item.Size,
			item.Total_price, item.Nm_id, item.Brand, item.Status).Scan(&lastInsertId)
	}
	if err != nil {
		log.Printf(": unable to insert data (items): %v\n", err)

	}

	ItemIdfk := lastInsertId
	err = tx.QueryRow(context.Background(), `INSERT INTO deliveries (delivery_id, name_, phone, zip, city, address_, region, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING delivery_id`,
		o.Delivery.Delivery_id, o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip, o.Delivery.City, o.Delivery.Address,
		o.Delivery.Region, o.Delivery.Email).Scan(&lastInsertId)
	if err != nil {
		log.Printf(": unable to insert data (deliveries): %v\n", err)

	}
	DeliveryIdfk := lastInsertId

	err = tx.QueryRow(context.Background(), `INSERT INTO payments (payment_id, transactions, request_id, currency, provider_, 
		amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING payment_id`,
		o.Payments.Payment_id, o.Payments.Transaction, o.Payments.Request_id, o.Payments.Currency, o.Payments.Provider,
		o.Payments.Amount, o.Payments.Payment_dt, o.Payments.Bank, o.Payments.Delivery_cost, o.Payments.Goods_total, o.Payments.Custom_fee).Scan(&lastInsertId)
	if err != nil {
		log.Printf(": unable to insert data (payments): %v\n", err)

	}
	PaymentId_fk := lastInsertId

	err = tx.QueryRow(context.Background(), `INSERT INTO orders (order_uid, track_number, name_entry, locale, internal_signature,
		customer_id, delivery_service, shardkey, sm_id, date_created, off_shard, fk_delivery_id, fk_payment_id, fk_item_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING order_uid`,
		o.Order_uid, o.Track_number, o.Entry, o.Locale, o.Internal_signature, o.Customer_id, o.Delivery_service, o.Shardkey,
		o.Sm_id, o.Date_created, o.Oof_shard, DeliveryIdfk, PaymentId_fk, ItemIdfk).Scan(&lastInsertOrderUid)
	if err != nil {
		log.Printf(": unable to insert data (orders): %v\n", err)

	}
	orderIdFk := lastInsertOrderUid

	err = tx.Commit(context.Background())
	if err != nil {
		return "0", err
	}
	fmt.Println(orderIdFk)

	S.csh.SetCache(orderIdFk, o)

	return orderIdFk, nil
}
func (S *DB) SetCacheFromDB() {
	var oid string

	err := S.db.QueryRow(context.Background(), `SELECT order_uid FROM orders`).Scan(&oid)
	if err != nil {
		log.Printf(": unable to get order_id from database: %v\n", err)
	}
	o, err := S.GetOrderFromDB(oid)
	if err != nil {
		log.Printf(": unable to get orders from database: %v\n", err)
	}
	S.csh.SetCache(oid, o)

}
func (S *DB) GetOrderFromDB(oid string) (storage.Orders, error) {
	var o storage.Orders
	var payment_id_fk, delivery_id_fk, item_id_fk int64

	err := S.db.QueryRow(context.Background(), `SELECT order_uid, track_number, name_entry, locale, internal_signature, customer_id, 
	delivery_service, shardkey, sm_id, date_created, off_shard, fk_delivery_id, fk_payment_id, fk_item_id 
	FROM orders WHERE order_uid = $1`,
		oid).Scan(&o.Order_uid, &o.Track_number, &o.Entry, &o.Locale, &o.Internal_signature, &o.Customer_id, &o.Delivery_service, &o.Shardkey,
		&o.Sm_id, &o.Date_created, &o.Oof_shard, &payment_id_fk, &delivery_id_fk, &item_id_fk)
	if err != nil {
		log.Printf(": unable to get orders from database: %v\n", err)
		return o, err
	}

	err = S.db.QueryRow(context.Background(), `SELECT payment_id, transactions, request_id, currency, provider_, amount, payment_dt,
	bank, delivery_cost, goods_total, custom_fee FROM payments WHERE payment_id = $1`,
		payment_id_fk).Scan(&o.Payments.Payment_id, &o.Payments.Transaction, &o.Payments.Request_id, &o.Payments.Currency,
		&o.Payments.Provider, &o.Payments.Amount, &o.Payments.Payment_dt, &o.Payments.Bank, &o.Payments.Delivery_cost,
		&o.Payments.Goods_total, &o.Payments.Custom_fee)
	if err != nil {
		log.Printf(": unable to get payment from database: %v\n", err)
		return o, err
	}

	err = S.db.QueryRow(context.Background(), `SELECT delivery_id, name_, phone, zip, city, address_, region, email
	FROM deliveries WHERE delivery_id = $1`,
		delivery_id_fk).Scan(&o.Delivery.Delivery_id, &o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, &o.Delivery.City,
		&o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email)
	if err != nil {
		log.Printf(": unable to get deliveries from database: %v\n", err)
		return o, err
	}
	for _, item := range o.Items {
		err = S.db.QueryRow(context.Background(), `SELECT chrt_id, track_number, price, rid, name_, sale, size_, total_price,
	nm_id, brand, status_ FROM items WHERE chrt_id = $1`,
			item_id_fk).Scan(&item.Chrt_id, &item.Track_number, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size,
			&item.Total_price, &item.Nm_id, &item.Brand, &item.Status)
		if err != nil {
			log.Printf(": unable to get items from database: %v\n", err)
			return o, err
		}
		o.Items = append(o.Items, item)
	}
	return o, nil
}
