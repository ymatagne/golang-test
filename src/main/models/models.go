package models

import "time"

// Event needing to be checked and verified
type Event struct {
	LoadID     string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	LoadAmount string    `json:"load_amount"`
	Time       time.Time `json:"time"`
}

// Account is the state of the Bank Account of the CustomerID
type Account struct {
	LoadID       string
	MaxPerDay    float64
	MaxPerWeek   float64
	NumberPerDay int
	LatestDate   time.Time
	History      map[string]bool
}

// Activity is the output of this command. Activity return the result of all the rules analysed anc verified
type Activity struct {
	LoadID     string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
