package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	payment_gateway_usecases "github.com/yescorihuela/deuna-payment-system/internal/application/usecases/payment_gateway"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)

type PaymentHandler struct {
	paymentUseCase payment_gateway_usecases.PaymentUseCase
}

func NewTransactionHandler(paymentUseCase payment_gateway_usecases.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

func (paymentHandler *PaymentHandler) Create(ctx *gin.Context) {
	req := requests.NewPaymentRequestRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entityTransaction, err := mappers.FromPaymentRequestToTransactionEntity(req)

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

