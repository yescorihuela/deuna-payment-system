package models

import "time"

type Refund struct {
	Id            string    `json:"id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
