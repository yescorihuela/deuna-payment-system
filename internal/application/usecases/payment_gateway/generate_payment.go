package usecases

import (
	"context"
	"net/http"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/constants"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/responses"
	http_client "github.com/yescorihuela/deuna-payment-system/internal/infrastructure/services/http"
)

type PaymentUseCase interface {
	Create(transaction entities.Transaction, creditCard entities.PaymentData) (*entities.Transaction, error)
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

func (uc *paymentUseCase) Create(transaction entities.Transaction, creditCard entities.PaymentData) (*entities.Transaction, error) {
	ctx := context.Background()
	tx, err := uc.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}
	// TODO: solve this
	var req = requests.PaymentRequest{
		Amount:          creditCard.Amount,
		Currency:        creditCard.Currency,
		CardNumber:      creditCard.CardNumber,
		ExpireDate:      creditCard.ExpireDate,
		CVV:             creditCard.CVV,
		MerchantCode:    creditCard.MerchantCode,
		TransactionType: creditCard.TransactionType,
	}

	res, err := uc.httpClient.Post(ctx, "transaction", req)

	if err != nil {
		uc.transactionRepository.SetTransactionStatus(transaction.MerchantCode, transaction.Id, constants.REJECTED)
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		uc.transactionRepository.SetTransactionStatus(transaction.MerchantCode, transaction.Id, constants.APPROVED)
		
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
