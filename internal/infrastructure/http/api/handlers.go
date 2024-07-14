package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application/usecases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)

type PaymentHandler struct {
	paymentUseCase usecases.PaymentUseCase
}

type RefundHandler struct {
	refundUseCase usecases.RefundUseCase
}

func NewTransactionHandler(paymentUseCase usecases.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

func NewRefundHandler(refundUseCase usecases.RefundUseCase) *RefundHandler {
	return &RefundHandler{
		refundUseCase: refundUseCase,
	}
}

func (paymentHandler *PaymentHandler) Create(ctx *gin.Context) {
	req := requests.NewPaymentRequestRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entityTransaction, err := mappers.FromPaymentRequestToTransactionEntity(req)

	// TODO: call http_client

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	savedTransaction, err := paymentHandler.paymentUseCase.Create(entityTransaction)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	
	paymentResponse := mappers.FromTransactionEntityToResponse(*savedTransaction)

	ctx.JSON(http.StatusCreated, paymentResponse)
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
