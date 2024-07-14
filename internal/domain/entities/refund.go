package entities

import "time"

type Refund struct {
	Id            string
	TransactionId string
	MerchantId    string
	Amount        float64
	Status        string
	CreatedAt     time.Time
}
