package usecases

import (
	"context"
	"errors"
	"net/http"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/constants"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/refund"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/repositories/transaction"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/responses"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
	http_client "github.com/yescorihuela/deuna-payment-system/internal/infrastructure/services/http"
)

type RefundUseCase interface {
	Create(transaction entities.Refund) (*entities.Transaction, error)
}

type refundUseCase struct {
	refundRepository      refund.RefundRepository
	transactionRepository transaction.TransactionRepository
	httpClient            http_client.HttpClientInterface[requests.RefundRequest, responses.RefundResponse]
}

func NewRefundUseCase(
	refundRepository refund.RefundRepository,
	transactionRepository transaction.TransactionRepository,
	httpClient http_client.HttpClientInterface[requests.RefundRequest, responses.RefundResponse],
) RefundUseCase {
	return &refundUseCase{
		refundRepository:      refundRepository,
		transactionRepository: transactionRepository,
		httpClient:            httpClient,
	}
}

func (uc *refundUseCase) Create(refund entities.Refund) (*entities.Transaction, error) {
	ctx := context.Background()

	transactionModel, err := uc.transactionRepository.GetPaymentByTransactionId(refund.MerchantId, refund.TransactionId)
	if err != nil {
		return nil, err
	}

	if refund.Amount > transactionModel.Amount {
		return nil, errors.New("refund amount never be greater than original transaction")
	}

	refundExists, err := uc.refundRepository.GetRefundByTransactionId(refund.MerchantId, refund.TransactionId)
	if err != nil {
		return nil, err
	}

	if refundExists {
		return nil, errors.New("refund process already exists")
	}

	_, err = uc.refundRepository.Create(refund)
	if err != nil {
		return nil, err
	}

	var req = requests.NewRefundRequest()

	req.Amount = refund.Amount
	req.TransactionId = refund.TransactionId
	req.MerchantId = refund.MerchantId
	req.TransactionType = constants.REFUND

	res, err := uc.httpClient.Post(ctx, "api/v1/transaction", req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		err := uc.transactionRepository.SetTransactionStatus(refund.MerchantId, refund.TransactionId, constants.REFUND)
		if err != nil {
			return nil, err
		}
	}

	transactionModel, err = uc.transactionRepository.GetPaymentByTransactionId(refund.MerchantId, refund.TransactionId)
	if err != nil {
		return nil, err
	}

	refundEntity := mappers.FromTransactionModelToEntity(*transactionModel)
	return &refundEntity, err
}
