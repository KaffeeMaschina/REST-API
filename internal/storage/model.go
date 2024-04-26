package storage

type Orders struct {
	Order_uid          string     `JSON:"order_uid"`
	Track_number       string     `JSON:"track_number"`
	Entry              string     `JSON:"entry"`
	Delivery           Deliveries `JSON:"deliveries"`
	Locale             string     `JSON:"locale"`
	Internal_signature string     `JSON:"internal_signature"`
	Customer_id        string     `JSON:"customer_id"`
	Delivery_service   string     `JSON:"delivery_service"`
	Shardkey           string     `JSON:"shardkey"`
	Sm_id              int32      `JSON:"sm_id"`
	Date_created       string     `JSON:"date_created"`
	Oof_shard          string     `JSON:"oof_shard"`
	Items              []Items    `JSON:"items"`
	Payments           Payments   `JSON:"payments"`
}

type Deliveries struct {
	Delivery_id int32  `JSON:"delivery_id"`
	Name        string `JSON:"name"`
	Phone       string `JSON:"phone"`
	Zip         string `JSON:"zip"`
	City        string `JSON:"city"`
	Address     string `JSON:"address"`
	Region      string `JSON:"region"`
	Email       string `JSON:"email"`
}

type Payments struct {
	Payment_id    int32  `JSON:"payment_id"`
	Transaction   string `JSON:"transactions"`
	Request_id    string `JSON:"request_id"`
	Currency      string `JSON:"currency"`
	Provider      string `JSON:"provider_"`
	Amount        int32  `JSON:"amount"`
	Payment_dt    int32  `JSON:"payment_dt"`
	Bank          string `JSON:"bank"`
	Delivery_cost int32  `JSON:"delivery_cost"`
	Goods_total   int32  `JSON:"goods_total"`
	Custom_fee    int32  `JSON:"custom_fee"`
}
type Items struct {
	Chrt_id      int32  `JSON:"chrt_id"`
	Track_number string `JSON:"track_number"`
	Price        int32  `JSON:"price"`
	Rid          string `JSON:"rid"`
	Name         string `JSON:"name_"`
	Sale         int32  `JSON:"sale"`
	Size         string `JSON:"size_"`
	Total_price  int32  `JSON:"total_price"`
	Nm_id        int32  `JSON:"nm_id"`
	Brand        string `JSON:"brand"`
	Status       int32  `JSON:"status_"`
}
