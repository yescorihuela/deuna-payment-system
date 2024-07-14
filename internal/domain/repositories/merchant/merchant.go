package merchant

import "github.com/yescorihuela/deuna-payment-system/internal/domain/entities"

type MerchantRepository interface {
	Create(merchant entities.Merchant) (*entities.Merchant, error)
	GetByMerchantCode(merchantCode string) (*entities.Merchant, error)
	GetById(id int) (*entities.Merchant, error)
	DisableMerchant(merchantCode string) error
	Update(id int, merchant entities.Merchant) (*entities.Merchant, error)
}
