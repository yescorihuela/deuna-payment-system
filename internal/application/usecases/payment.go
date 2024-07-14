package usecases

import (
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
)

type TransactionUseCase interface {
	Create(transaction entities.Transaction) (*entities.Transaction, error)
}

type transactionUseCase struct {
	transactionRepository transaction.TransactionRepository
}

func NewTransaction(transactionRepository transaction.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		transactionRepository,
	}
}

func (uc *transactionUseCase) Create(transaction entities.Transaction) (*entities.Transaction, error) {
	tx, err := uc.transactionRepository.Create(transaction)
	return tx, err
}
