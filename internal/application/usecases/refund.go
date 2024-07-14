package usecases

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/repositories"
)

type RefundUseCase interface {
	Create(transaction entities.Refund) (*entities.Refund, error)
}

type refundUseCase struct {
	refundRepository repositories.PostgresqlRefundRepository
}

func NewRefund(refundRepository repositories.PostgresqlRefundRepository) RefundUseCase {
	return &refundUseCase{
		refundRepository,
	}
}

func (uc *refundUseCase) Create(refund entities.Refund) (*entities.Refund, error) {
	ref, err := uc.refundRepository.Create(refund)
	return ref, err
}
