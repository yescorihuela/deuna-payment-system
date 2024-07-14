package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application/usecases"
)

type TransactionHandler struct {
	transactionUseCase usecases.TransactionUseCase
}

func NewTransactionHandler(transaction usecases.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		transactionUseCase: transaction,
	}
}

func (app *TransactionHandler) Create(ctx *gin.Context) {

}
