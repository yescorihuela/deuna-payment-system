package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application/payment_gateway"
	"github.com/yescorihuela/deuna-payment-system/internal/application/usecases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/databases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/responses"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/repositories"
	http_client "github.com/yescorihuela/deuna-payment-system/internal/infrastructure/services/http"
)

func main() {
	db, err := databases.NewPostgresqlDbConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pgTransactionRepository := repositories.NewPostgresqlTransactionRepository(db)

	pgRefundRepository := repositories.NewPostgresqlRefundRepository(db)
	httpClient := http_client.NewHttpClient[requests.PaymentRequest, responses.PaymentResponse](http_client.HttpClientSettings{})
	paymentProcessUseCase := usecases.NewPaymentProcess(
		pgTransactionRepository,
		httpClient,
	)
	refundUseCase := usecases.NewRefundUseCase(pgRefundRepository)

	txHandler := api.NewTransactionHandler(paymentProcessUseCase)
	refundHandler := api.NewRefundHandler(refundUseCase)

	txApp := payment_gateway.NewApplication(
		txHandler,
		refundHandler,
		gin.Default())

	if err := txApp.Run(); err != nil {
		panic(err)
	}
}
