package refund

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
)

type RefundRepository interface {
	Create(refund entities.Refund) (*models.Refund, error)
}
