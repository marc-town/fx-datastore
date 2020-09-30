package model

import "time"

type Transactions []Transaction

type Transaction struct {
	ID           int
	CurrencyPair string  `json:"currency_pair"`
	Direction    int     `json:"direction"`
	EntryPoint   float64 `json:"entry_point"`
	EntryTime    string  `json:"entry_time"`
	ExitPoint    float64 `json:"exit_point"`
	ExitTime     string  `json:"exit_time"`
	Lot          float64 `json:"lot"`
	Spread       float64 `json:"spread"`
	Commission   int     `json:"commission"`
	Swap         int     `json:"swap"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
