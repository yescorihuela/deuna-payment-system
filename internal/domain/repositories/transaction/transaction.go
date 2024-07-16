package transaction

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
)

type TransactionRepository interface {
	Create(refund entities.Transaction) (*models.Transaction, error)
	SetTransactionStatus(merchantCode, transactionId, status string) error
	GetPaymentByTransactionId(merchantCode, transactionId string) (*models.Transaction, error)
	GetAllTransactionsByMerchant(merchantCode string) ([]*models.Transaction, error)
}
