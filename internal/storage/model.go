package storage

type Order struct {
	Order_uid          string  `JSON:"order_uid"`
	Track_number       string  `JSON:"track_number"`
	Entry              string  `JSON:"entry"`
	Delivery           string  `JSON:"name"`
	Locale             string  `JSON:"locale"`
	Internal_signature string  `JSON:"internal_signature"`
	Customer_id        string  `JSON:"customer_id"`
	Delivery_service   string  `JSON:"delivery_service"`
	Shardkey           int     `JSON:"shardkey"`
	Sm_id              int     `JSON:"sm_id"`
	Date_created       int     `JSON:"date_created"`
	Oof_shard          int     `JSON:"oof_shard"`
	Items              []Items `JSON:"items"`
	Payment            Payment `JSON:"payment"`
}

type Delivery struct {
	Name    string `JSON:"name"`
	Phone   string `JSON:"phone"`
	Zip     int    `JSON:"zip"`
	City    string `JSON:"city"`
	Address string `JSON:"address"`
	Region  string `JSON:"region"`
	Email   string `JSON:"email"`
}

type Payment struct {
	Transaction   string `JSON:"transaction"`
	Request_id    string `JSON:"request_id"`
	Currency      string `JSON:"currency"`
	Provider      string `JSON:"provider"`
	Amount        int    `JSON:"amount"`
	Payment_dt    int    `JSON:"payment_dt"`
	Bank          string `JSON:"bank"`
	Delivery_cost int    `JSON:"delivery_cost"`
	Goods_total   int    `JSON:"goods_total"`
	Custom_fee    int    `JSON:"custom_fee"`
}
type Items struct {
	Chrt_id      int    `JSON:"chrt_id"`
	Track_number string `JSON:"track_number"`
	Price        int    `JSON:"price"`
	Rid          string `JSON:"rid"`
	Name         string `JSON:"name"`
	Sale         int    `JSON:"sale"`
	Size         int    `JSON:"size"`
	Total_price  int    `JSON:"total_price"`
	Nm_id        int    `JSON:"nm_id"`
	Brand        string `JSON:"brand"`
	Status       int    `JSON:"status"`
}
