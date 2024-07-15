package handlers


import (
	"net/http"

	"github.com/gin-gonic/gin"
	payment_gateway_usecases "github.com/yescorihuela/deuna-payment-system/internal/application/usecases/payment_gateway"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)
type RefundHandler struct {
	refundUseCase payment_gateway_usecases.RefundUseCase
}
func NewRefundHandler(refundUseCase payment_gateway_usecases.RefundUseCase) *RefundHandler {
	return &RefundHandler{
		refundUseCase: refundUseCase,
	}
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
	responseRefund := mappers.FromRefundEntityToResponse(*savedRefund)
	ctx.JSON(http.StatusCreated, responseRefund)
}
