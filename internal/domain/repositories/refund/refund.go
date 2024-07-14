package refund

import "github.com/yescorihuela/deuna-payment-system/internal/domain/entities"

type RefundRepository interface {
	Create(refund entities.Refund) (*entities.Refund, error)
}
