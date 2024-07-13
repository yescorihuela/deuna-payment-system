package entities

import "time"

type Transaction struct {
	Id         string
	MerchantId string
	Amount     float64
	Status     string
	CreatedAt  time.Time
}
