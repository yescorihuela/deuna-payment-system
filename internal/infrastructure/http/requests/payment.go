package requests

type PaymentRequest struct {
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	CardNumber      string  `json:"card_number"`
	ExpireDate      string  `json:"expire_date"`
	CVV             string  `json:"cvv"`
	MerchantCode    string  `json:"merchant_code"`
	TransactionType string  `json:"transaction_type,omitempty"`
}

func NewPaymentRequest() PaymentRequest {
	return PaymentRequest{}
}
