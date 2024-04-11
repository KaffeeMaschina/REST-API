CREATE TABLE IF NOT EXIST PAYMENT 
(
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
);

CREATE TABLE ITEMS 
(
chrt_id int primary key not null,
track_number varchar(128),
price int,
rid varchar(128),
"name" varchar(128),
sale int,
"size" int,
total_price int,
nm_id int,
brand varchar(128),
"status" int
);

CREATE TABLE DELIVERY
(
delivery_id int primary key not null,
"name" varchar(128),
phone varchar(128),
zip int,
city varchar(128),
"address" varchar(128),
region varchar(128),
email varchar(128)
);

CREATE TABLE "ORDER"
(
order_uid int primary key not null,
track_number varchar(128),
"entry" varchar(128),
locale varchar(128),
internal_signature varchar(128),
customer_id varchar(128),
delivery_service varchar(128),
shardkey int,
sm_id int,
date_created timestamp,
off_shard int,
fk_delivery_id int references delivery(delivery_id) not null,
fk_payment_id int references payment(payment_id) not null,
fk_item_id int references items(chrt_id) not null
);