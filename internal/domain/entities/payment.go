package entities

type PaymentData struct {
	Amount          float64
	Currency        string
	CardNumber      string
	ExpireDate      string
	CVV             string
	MerchantCode    string
	TransactionType string
}
