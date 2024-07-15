package merchant

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
)

type MerchantRepository interface {
	Create(merchant entities.Merchant) (*models.Merchant, error)
	GetByMerchantCode(merchantCode string) (*models.Merchant, error)
	GetById(id int) (*models.Merchant, error)
	SetStatus(merchantCode string, isEnabled bool) error
	Update(merchantCode string, merchant entities.Merchant) (*models.Merchant, error)
	ExecuteTransaction(merchantCode, kindTransaction string, amount float64) error
}
