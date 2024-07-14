package usecases

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/refund"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)

type RefundUseCase interface {
	Create(transaction entities.Refund) (*entities.Refund, error)
}

type refundUseCase struct {
	refundRepository refund.RefundRepository
}

func NewRefundUseCase(refundRepository refund.RefundRepository) RefundUseCase {
	return &refundUseCase{
		refundRepository,
	}
}

func (uc *refundUseCase) Create(refund entities.Refund) (*entities.Refund, error) {
	refundModel, err := uc.refundRepository.Create(refund)
	refundEntity := mappers.FromRefundModelToEntity(*refundModel)
	return &refundEntity, err
}
