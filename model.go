package restapi

type order struct {
	order_uid    string `JSON:"order_uid"`
	track_number string `JSON:"track_number"`
	entry        string `JSON:"entry"`
	delivery     string `JSON:"name"`
}

type payment struct {
	transaction string `JSON:"transaction"`
}
