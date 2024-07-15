package acquiring_bank

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/merchant"
)

type MerchantUseCase interface {
	Create(merchant entities.Merchant) (*entities.Merchant, error)
	Update(merchantCode string, merchant entities.Merchant) (*entities.Merchant, error)
	SetStatus(merchantCode string, isEnabled bool) error
	GetByMerchantCode(merchantCode string) (*entities.Merchant, error)
	GetById(id string) (*entities.Merchant, error)
}

type merchantUseCase struct {
	merchantRepository merchant.MerchantRepository
}

func NewMerchantUseCase(merchantRepository merchant.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{
		merchantRepository: merchantRepository,
	}
}

func (uc *merchantUseCase) Create(merchant entities.Merchant) (*entities.Merchant, error) {
	return nil, nil
}

func (uc *merchantUseCase) Update(merchantCode string, merchant entities.Merchant) (*entities.Merchant, error) {
	return nil, nil
}

func (uc *merchantUseCase) SetStatus(merchantCode string, isEnabled bool) error {
	return nil
}

func (uc *merchantUseCase) GetByMerchantCode(merchantCode string) (*entities.Merchant, error) {
	return nil, nil
}

func (uc *merchantUseCase) GetById(id string) (*entities.Merchant, error) {
	return nil, nil
}
