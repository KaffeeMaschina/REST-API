package storage

import (
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

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

func (o Orders) ValidateOrders() error {
	fmt.Println("validation")
	return validation.ValidateStruct(
		&o,
		validation.Field(&o.Order_uid, AlphanumericRule...),
		validation.Field(&o.Track_number, is.UpperCase, is.Alpha),
		validation.Field(&o.Entry, validation.Required, is.UpperCase, is.Alpha),
		validation.Field(&o.Locale, validation.Required, validation.Match(regexp.MustCompile("^[a-z]{2}$"))),
		validation.Field(&o.Internal_signature),
		validation.Field(&o.Customer_id, AlphaRule...),
		validation.Field(&o.Delivery_service, AlphaRule...),
		validation.Field(&o.Shardkey, DigitRule...),
		validation.Field(&o.Sm_id, validation.Required),
		validation.Field(&o.Date_created, validation.Required, validation.Date(time.RFC3339)),
		validation.Field(&o.Oof_shard, DigitRule...),
		//Delivery validation
		validation.Field(&o.Delivery),
		//Items validation
		validation.Field(&o.Items),
		//Payments validation
		validation.Field(&o.Payments),
	)
}

// Delivery validation
func (d Deliveries) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Delivery_id),
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Phone, validation.Required),
		validation.Field(&d.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{7}$"))),
		validation.Field(&d.City, validation.Required),
		validation.Field(&d.Address, validation.Required),
		validation.Field(&d.Region, AlphaRule...),
		validation.Field(&d.Email, validation.Required, is.Email),
	)
}

// Item validation
func (i Items) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Chrt_id, validation.Required),
		validation.Field(&i.Track_number, is.UpperCase, is.Alpha),
		validation.Field(&i.Price, validation.Required),
		validation.Field(&i.Rid, AlphanumericRule...),
		validation.Field(&i.Name, AlphaRule...),
		validation.Field(&i.Sale),
		validation.Field(&i.Size, validation.Required),
		validation.Field(&i.Total_price, validation.Required),
		validation.Field(&i.Nm_id, validation.Required),
		validation.Field(&i.Brand),
		validation.Field(&i.Status, validation.Required),
	)
}

// Payment validation
func (p Payments) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Payment_id),
		validation.Field(&p.Transaction, AlphanumericRule...),
		validation.Field(&p.Request_id),
		validation.Field(&p.Currency, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{3}$"))),
		validation.Field(&p.Provider, AlphaRule...),
		validation.Field(&p.Amount, validation.Required),
		validation.Field(&p.Payment_dt, validation.Required),
		validation.Field(&p.Bank, AlphaRule...),
		validation.Field(&p.Delivery_cost, validation.Required),
		validation.Field(&p.Goods_total, validation.Required),
		validation.Field(&p.Custom_fee),
	)
}

var AlphanumericRule = []validation.Rule{
	validation.Required,
	is.LowerCase,
	is.Alphanumeric,
}

var DigitRule = []validation.Rule{
	validation.Required,
	is.Digit,
}
var AlphaRule = []validation.Rule{
	validation.Required,
	is.Alpha,
}
