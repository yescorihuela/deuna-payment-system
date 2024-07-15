package responses

import "time"

type MerchantResponse struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Balance           float64   `json:"balance"`
	NotificationEmail string    `json:"notification_email"`
	MerchantCode      string    `json:"merchant_code"`
	Enabled           bool      `json:"enabled"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func NewMerchantResponse() MerchantResponse {
	return MerchantResponse{}
}
