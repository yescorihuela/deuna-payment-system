package models

import "time"

type Transaction struct {
	Id         string    `json:"string"`
	MerchantId string    `json:"merchant_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
