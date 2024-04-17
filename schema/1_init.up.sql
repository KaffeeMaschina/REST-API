
CREATE TABLE IS EXISTS payment 
(
payment_id INT PRIMARY KEY NOT NULL,
transactions  VARCHAR(128),
request_id VARCHAR(128),
currency VARCHAR(128),
provider VARCHAR(128),
amount INT,
payment_db INT,
bank VARCHAR(128),
delivery_cost INT,
goods_total INT,
custom_fee INT
);
/*
CREATE TABLE items 
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

CREATE TABLE delivery
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

CREATE TABLE "order"
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
fk_delivery_id int not null,
fk_payment_id int not null,
fk_item_id int not null
);*/