package mappers

import (
	"time"

	"github.com/yescorihuela/deuna-payment-system/internal/domain/constants"
	"github.com/yescorihuela/deuna-payment-system/internal/domain/entities"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/responses"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/models"
)

func FromPaymentRequestToTransactionEntity(request requests.PaymentRequest) entities.Transaction {

	// TODO: Validate credit card number
	newTxId := entities.NewUlid()
	txNow := time.Now().UTC()
	txEntity := entities.Transaction{
		Id:           newTxId,
		MerchantCode: request.MerchantCode,
		Amount:       request.Amount,
		Status:       constants.PENDING,
		CreatedAt:    txNow,
	}
	return txEntity
}

func FromRefundRequestToRefundEntity(request requests.RefundRequest) entities.Refund {
	return entities.Refund{}
}

func FromPaymentRequestToPaymentEntity(request requests.PaymentRequest) entities.PaymentData {
	return entities.PaymentData{
		Amount:          request.Amount,
		Currency:        request.Currency,
		CardNumber:      request.CardNumber,
		ExpireDate:      request.ExpireDate,
		CVV:             request.CVV,
		MerchantCode:    request.MerchantCode,
		TransactionType: request.TransactionType,
	}
}

func FromRefundModelToEntity(refund models.Refund) entities.Refund {
	return entities.Refund{
		Id:            refund.Id,
		TransactionId: refund.TransactionId,
		MerchantId:    refund.MerchantId,
		Amount:        refund.Amount,
		Status:        refund.Status,
		CreatedAt:     refund.CreatedAt,
	}
}

func FromRefundEntityToModel(refund entities.Refund) models.Refund {
	return models.Refund{
		Id:            refund.Id,
		TransactionId: refund.TransactionId,
		MerchantId:    refund.MerchantId,
		Amount:        refund.Amount,
		Status:        refund.Status,
		CreatedAt:     time.Now(),
	}
}

func FromMerchantRequestToEntity(merchant requests.MerchantRequest) entities.Merchant {
	now := time.Now().UTC()
	merchantCode := entities.NewNanoId()
	return entities.Merchant{
		Name:              merchant.Name,
		Balance:           merchant.Balance,
		NotificationEmail: merchant.NotificationEmail,
		MerchantCode:      merchantCode,
		Enabled:           merchant.Enabled,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

func FromMerchantRequestToEntityUpdate(merchant requests.MerchantRequest) entities.Merchant {
	now := time.Now().UTC()
	return entities.Merchant{
		Name:              merchant.Name,
		Balance:           merchant.Balance,
		NotificationEmail: merchant.NotificationEmail,
		MerchantCode:      merchant.MerchantCode,
		Enabled:           merchant.Enabled,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

func FromMerchantEntityToModel(merchant entities.Merchant) models.Merchant {
	return models.Merchant{
		Id:                merchant.Id,
		Name:              merchant.Name,
		Balance:           merchant.Balance,
		NotificationEmail: merchant.NotificationEmail,
		MerchantCode:      merchant.MerchantCode,
		Enabled:           merchant.Enabled,
	}
}

func FromMerchantModelToEntity(merchant models.Merchant) entities.Merchant {
	return entities.Merchant{
		Id:                merchant.Id,
		Name:              merchant.Name,
		Balance:           merchant.Balance,
		NotificationEmail: merchant.NotificationEmail,
		MerchantCode:      merchant.MerchantCode,
		Enabled:           merchant.Enabled,
	}
}

func FromTransactionEntityToResponse(transaction entities.Transaction) responses.PaymentResponse {
	return responses.PaymentResponse{
		Id:        transaction.Id,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}
}

func FromRefundEntityToResponse(refund entities.Refund) responses.RefundResponse {
	return responses.RefundResponse{}
}

func FromMerchantEntityToResponse(merchant entities.Merchant) responses.MerchantResponse {
	return responses.MerchantResponse{
		Id:                merchant.Id,
		Name:              merchant.Name,
		Balance:           merchant.Balance,
		NotificationEmail: merchant.NotificationEmail,
		MerchantCode:      merchant.MerchantCode,
		Enabled:           merchant.Enabled,
		CreatedAt:         merchant.CreatedAt,
		UpdatedAt:         merchant.UpdatedAt,
	}
}

func FromTransactionModelToEntity(transaction models.Transaction) entities.Transaction {
	return entities.Transaction{
		Id:           transaction.Id,
		MerchantCode: transaction.MerchantId,
		Amount:       transaction.Amount,
		Status:       transaction.Status,
		CreatedAt:    transaction.CreatedAt,
	}
}
