package models
import "time"

type Attempt struct {
	Id   string `json:"id"`
	CustomerId   string `json:"customer_id"`
	LoadAmount    string    `json:"load_amount"`
	Time	time.Time `json:"time"`
}