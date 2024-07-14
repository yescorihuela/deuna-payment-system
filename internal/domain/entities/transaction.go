package entities

import (
	"time"

	"github.com/oklog/ulid/v2"
)

func NewUlid() string {
	return ulid.Make().String()
}

type Transaction struct {
	Id           string
	MerchantCode string
	Amount       float64
	Status       string
	CreatedAt    time.Time
}
