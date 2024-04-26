BEGIN;

CREATE TABLE IF NOT EXISTS payments (
payment_id INT PRIMARY KEY,
transactions  VARCHAR(128),
request_id VARCHAR(128),
currency VARCHAR(128),
provider_ VARCHAR(128),
amount INT,
payment_dt INT,
bank VARCHAR(128),
delivery_cost INT,
goods_total INT,
custom_fee INT
);

CREATE TABLE IF NOT EXISTS items (
chrt_id INT PRIMARY KEY,
track_number VARCHAR(128),
price INT,
rid VARCHAR(128),
name_ VARCHAR(128),
sale INT,
size_ VARCHAR(128),
total_price INT,
nm_id INT,
brand VARCHAR(128),
status_ INT
);

CREATE TABLE IF NOT EXISTS deliveries (
delivery_id INT PRIMARY KEY NOT NULL,
name_ VARCHAR(128),
phone VARCHAR(128),
zip VARCHAR(128),
city VARCHAR(128),
address_ VARCHAR(128),
region VARCHAR(128),
email VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS orders (
order_uid VARCHAR(128) PRIMARY KEY NOT NULL,
track_number VARCHAR(128),
name_entry VARCHAR(128),
locale VARCHAR(128),
internal_signature VARCHAR(128),
customer_id VARCHAR(128),
delivery_service VARCHAR(128),
shardkey VARCHAR(128),
sm_id INT,
date_created VARCHAR(128),
off_shard VARCHAR(128),
fk_delivery_id INT NOT NULL,
fk_payment_id INT NOT NULL,
fk_item_id INT NOT NULL
);

ALTER TABLE orders
ADD CONSTRAINT fk_order_delivery
FOREIGN KEY (fk_delivery_id) REFERENCES deliveries(delivery_id);

ALTER TABLE orders
ADD CONSTRAINT fk_order_payment
FOREIGN KEY (fk_payment_id) REFERENCES payments(payment_id);

ALTER TABLE orders
ADD CONSTRAINT fk_order_item
FOREIGN KEY (fk_item_id) REFERENCES items(chrt_id);

COMMIT;