package models

import "time"

type Attempt struct {
	Id         string    `json:"id"`
	CustomerId string    `json:"customer_id"`
	LoadAmount string    `json:"load_amount"`
	Time       time.Time `json:"time"`
}

type Account struct {
	Id           string
	MaxPerDay    float64
	MaxPerWeek   float64
	NumberPerDay int
	LatestDate   time.Time `json:"time"`
}
type Activity struct {
	Id         string `json:"id"`
	CustomerId string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
