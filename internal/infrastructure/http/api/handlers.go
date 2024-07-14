package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application/usecases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)

type TransactionHandler struct {
	transactionUseCase usecases.TransactionUseCase
}

type RefundHandler struct {
	refundUseCase usecases.RefundUseCase
}

func NewTransactionHandler(transactionUseCase usecases.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionUseCase: transactionUseCase,
	}
}

func NewRefundHandler(refundUseCase usecases.RefundUseCase) *RefundHandler {
	return &RefundHandler{
		refundUseCase: refundUseCase,
	}
}

func (txHandler *TransactionHandler) Create(ctx *gin.Context) {
	req := requests.NewCreditCardRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entityTransaction, err := mappers.FromCreditCardRequestToTransactionEntity(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	savedTransaction, err := txHandler.transactionUseCase.Create(entityTransaction)
	// TODO: transactionResponse => mappers
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, savedTransaction)
}

func (refundHandler *RefundHandler) Create(ctx *gin.Context) {
	req := requests.NewRefundRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entityRefund, err := mappers.FromRefundRequestToRefundEntity(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err})
		return
	}

	savedRefund, err := refundHandler.refundUseCase.Create(entityRefund)
	// TODO: RefundResponse => mappers
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, savedRefund)
}
