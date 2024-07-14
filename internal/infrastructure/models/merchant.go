package models

import "time"

type Merchant struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Balance           float64   `json:"balance"`
	NotificationEmail string    `json:"notification_email"`
	MerchantCode      string    `json:"merchant_code"`
	Enabled           bool      `json:"enabled"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func NewMerchantModel() Merchant {
	return Merchant{}
}
