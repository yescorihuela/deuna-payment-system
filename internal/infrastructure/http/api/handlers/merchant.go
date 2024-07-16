package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	acquiring_bank_usecases "github.com/yescorihuela/deuna-payment-system/internal/application/usecases/acquiring_bank"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/mappers"
)

type AcquiringBankHandler struct {
	acquiringBankUseCase acquiring_bank_usecases.MerchantUseCase
}

func NewAcquiringBankHandler(acquiringBankUseCase acquiring_bank_usecases.MerchantUseCase) *AcquiringBankHandler {
	return &AcquiringBankHandler{
		acquiringBankUseCase,
	}
}

func (acquiringBankHandler *AcquiringBankHandler) New(ctx *gin.Context) {
	req := requests.NewMerchantRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchantEntity := mappers.FromMerchantRequestToEntity(req)

	savedMerchant, err := acquiringBankHandler.acquiringBankUseCase.Create(merchantEntity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	merchantResponse := mappers.FromMerchantEntityToResponse(*savedMerchant)
	ctx.JSON(http.StatusCreated, merchantResponse)

}

func (acquiringBankHandler *AcquiringBankHandler) Update(ctx *gin.Context) {

	req := requests.NewMerchantRequest()
	merchantCodeParam := ctx.Param("merchant_code")
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchantEntity := mappers.FromMerchantRequestToEntityUpdate(req)

	updatedMerchant, err := acquiringBankHandler.acquiringBankUseCase.Update(merchantCodeParam, merchantEntity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	merchantResponse := mappers.FromMerchantEntityToResponse(*updatedMerchant)

	ctx.JSON(http.StatusOK, merchantResponse)

}

func (acquiringBankHandler *AcquiringBankHandler) ChangeStatus(ctx *gin.Context) {
	merchantCodeParam := ctx.Param("merchant_code")
	req := struct {
		Status bool `json:"status"`
	}{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := acquiringBankHandler.acquiringBankUseCase.SetStatus(merchantCodeParam, req.Status)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (acquiringBankHandler *AcquiringBankHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if strings.TrimSpace(idParam) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "merchant_code param in blank"})
		return
	}

	merchantEntity, err := acquiringBankHandler.acquiringBankUseCase.GetById(idParam)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	merchantResponse := mappers.FromMerchantEntityToResponse(*merchantEntity)

	ctx.JSON(http.StatusOK, merchantResponse)
}

func (acquiringBankHandler *AcquiringBankHandler) GetByMerchantCode(ctx *gin.Context) {
	merchantCodeParam := ctx.Param("merchant_code")
	if strings.TrimSpace(merchantCodeParam) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "merchant_code param in blank"})
		return
	}

	merchantEntity, err := acquiringBankHandler.acquiringBankUseCase.GetByMerchantCode(merchantCodeParam)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	merchantResponse := mappers.FromMerchantEntityToResponse(*merchantEntity)

	ctx.JSON(http.StatusOK, merchantResponse)
}

func (acquiringBankHandler *AcquiringBankHandler) ExecuteTransaction(ctx *gin.Context) {
	req := requests.NewPaymentRequest()
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.TransactionType == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("transaction type in blank")})
		return
	}
	err := acquiringBankHandler.acquiringBankUseCase.ExecuteTransaction(req.MerchantCode, req.TransactionType, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
