package responses

import "time"

type PaymentResponse struct {
	Id        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPaymentResponse() PaymentResponse {
	return PaymentResponse{}
}
