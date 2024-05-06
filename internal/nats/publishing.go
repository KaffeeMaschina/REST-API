package nats

import (
	"encoding/json"
	"log"
	"time"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage"
	"github.com/nats-io/stan.go"
)

type Publisher struct {
	sc *stan.Conn
}

// Create publisher
func NewPublisher(conn *stan.Conn) *Publisher {
	return &Publisher{
		sc: conn,
	}
}

// Publish to channel some data
func (p *Publisher) Publish() {
	item := storage.Items{Chrt_id: 9934930, Track_number: "WBILMTESTTRACK", Price: 453, Rid: "ab4219087a764ae0btest", Name: "Mascaras",
		Sale: 30, Size: "0", Total_price: 317, Nm_id: 2389212, Brand: "Vivienne Sabo", Status: 202}

	delivery := storage.Deliveries{Name: "Test Testov", Phone: "+9720000000", Zip: "2639809", City: "Kiryat Mozkin",
		Address: "Ploshad Mira 15", Region: "Kraiot", Email: "test@gmail.com"}

	payment := storage.Payments{Transaction: "b563feb7b2b84b6test", Request_id: "", Currency: "USD", Provider: "wbpay", Amount: 1817,
		Payment_dt: 1637907727, Bank: "alpha", Delivery_cost: 1500, Goods_total: 317, Custom_fee: 0}

	order := storage.Orders{Order_uid: "b563feb7b2b84b6test", Track_number: "WBILMTESTTRACK", Entry: "WBIL", Delivery: delivery,
		Locale: "en", Internal_signature: "", Payments: payment, Items: []storage.Items{item}, Customer_id: "test",
		Delivery_service: "meest", Shardkey: "9", Sm_id: 99, Date_created: "2021-11-26T06:22:19Z", Oof_shard: "1"}

	orderData, err := json.Marshal(order)
	if err != nil {
		log.Printf("Json.Marshal error:%s\n", err)
	}
	err = (*p.sc).Publish("Order", orderData)
	if err != nil {
		log.Printf("error publishing msg %s: \n", err)
	}
	time.Sleep(2 * time.Second)

}
