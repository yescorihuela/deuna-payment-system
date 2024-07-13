package usecases

type Refund interface{}

type refund struct{}

func NewRefund() Refund {
	return &refund{}
}
