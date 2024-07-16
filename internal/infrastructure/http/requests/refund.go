package requests

type RefundRequest struct {
	TransactionId   string  `json:"transaction_id"`
	MerchantId      string  `json:"merchant_code"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type,omitempty"`
}

func NewRefundRequest() RefundRequest {
	return RefundRequest{}
}
