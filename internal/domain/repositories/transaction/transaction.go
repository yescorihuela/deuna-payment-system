package transaction

import "github.com/yescorihuela/deuna-payment-system/internal/domain/entities"

type TransactionRepository interface {
	Create(refund entities.Transaction) (*entities.Transaction, error)
}
