package usecases

import (
	"context"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/constants"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/responses"
	http_client "github.com/yescorihuela/deuna-payment-system/internal/infrastructure/services/http"
)

type PaymentUseCase interface {
	Create(transaction entities.Transaction) (*entities.Transaction, error)
	SetPaymentStatus(merchantCode, transactionId, status string) error
}

type paymentUseCase struct {
	transactionRepository transaction.TransactionRepository
	httpClient            http_client.HttpClientInterface[requests.PaymentRequest, responses.PaymentResponse]
}

func NewPaymentProcess(
	transactionRepository transaction.TransactionRepository,
	httpClient http_client.HttpClientInterface[requests.PaymentRequest, responses.PaymentResponse],
) PaymentUseCase {
	return &paymentUseCase{
		transactionRepository: transactionRepository,
		httpClient:            httpClient,
	}
}

func (uc *paymentUseCase) Create(transaction entities.Transaction) (*entities.Transaction, error) {
	ctx := context.Background()
	tx, err := uc.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}
	// TODO: solve this
	var req requests.PaymentRequest
	_, err = uc.httpClient.Post(ctx, "payment-process", req)
	if err != nil {
		uc.transactionRepository.SetTransactionStatus(transaction.MerchantCode, transaction.Id, constants.REJECTED)
		return nil, err
	}
	return tx, err
}

func (uc *paymentUseCase) SetPaymentStatus(merchantCode, transactionId, status string) error {
	err := uc.transactionRepository.SetTransactionStatus(merchantCode, transactionId, status)
	if err != nil {
		return err
	}
	return nil
}
