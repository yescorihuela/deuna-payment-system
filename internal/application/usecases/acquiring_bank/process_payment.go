package acquiring_bank

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/merchant"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)

type MerchantUseCase interface {
	Create(merchant entities.Merchant) (*entities.Merchant, error)
	Update(merchantCode string, merchant entities.Merchant) (*entities.Merchant, error)
	SetStatus(merchantCode string, isEnabled bool) error
	GetByMerchantCode(merchantCode string) (*entities.Merchant, error)
	GetById(id string) (*entities.Merchant, error)
	ExecuteTransaction(merchantCode, transactionType string, amount float64) error
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
	merchantModel, err := uc.merchantRepository.Create(merchant)
	if err != nil {
		return nil, err
	}

	merchantEntity := mappers.FromMerchantModelToEntity(*merchantModel)

	return &merchantEntity, nil
}

func (uc *merchantUseCase) Update(merchantCode string, merchant entities.Merchant) (*entities.Merchant, error) {
	merchantModel, err := uc.merchantRepository.Update(merchantCode, merchant)
	if err != nil {
		return nil, err
	}
	merchantEntity := mappers.FromMerchantModelToEntity(*merchantModel)

	return &merchantEntity, nil
}

func (uc *merchantUseCase) SetStatus(merchantCode string, isEnabled bool) error {
	err := uc.merchantRepository.SetStatus(merchantCode, isEnabled)
	if err != nil {
		return err
	}

	return nil
}

func (uc *merchantUseCase) GetByMerchantCode(merchantCode string) (*entities.Merchant, error) {
	merchantModel, err := uc.merchantRepository.GetByMerchantCode(merchantCode)
	if err != nil {
		return nil, err
	}
	merchantEntity := mappers.FromMerchantModelToEntity(*merchantModel)
	return &merchantEntity, nil
}

func (uc *merchantUseCase) GetById(id string) (*entities.Merchant, error) {
	merchantModel, err := uc.merchantRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	merchantEntity := mappers.FromMerchantModelToEntity(*merchantModel)
	return &merchantEntity, nil
}

func (uc *merchantUseCase) ExecuteTransaction(merchantCode, transactionType string, amount float64) error {
	return uc.merchantRepository.ExecuteTransaction(merchantCode, transactionType, amount)
}
