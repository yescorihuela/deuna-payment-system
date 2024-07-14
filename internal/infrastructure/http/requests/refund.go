package requests

type RefundRequest struct {
	TransactionId string  `json:"transaction_id"`
	MerchantId    string  `json:"merchant_id"`
	Amount        float64 `json:"amount"`
}

func NewRefundRequest() RefundRequest {
	return RefundRequest{}
}
