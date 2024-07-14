package models

type Merchant struct {
	Id                int     `json:"Id"`
	Name              string  `json:"name"`
	Balance           float64 `json:"balance"`
	NotificationEmail string  `json:"notification_email"`
	MerchantCode      string  `json:"merchant_code"`
}
