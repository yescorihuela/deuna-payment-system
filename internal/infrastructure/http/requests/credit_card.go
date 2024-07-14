package requests

type CreditCardRequest struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	CardNumber string  `json:"card_number"`
	ExpireDate string  `json:"expire_date"`
	CVV        string  `json:"cvv"`
}

func NewCreditCardRequest() CreditCardRequest {
	return CreditCardRequest{}
}
