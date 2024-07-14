package requests

import "time"

type MerchantRequest struct {
	Id                string    `json:"id,omitempty"`
	Name              string    `json:"name"`
	Balance           float64   `json:"balance"`
	NotificationEmail string    `json:"notification_email"`
	MerchantCode      string    `json:"merchant_code"`
	Enabled           bool      `json:"enabled"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

func NewMerchantRequest() MerchantRequest {
	return MerchantRequest{}
}
