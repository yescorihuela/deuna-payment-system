package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application/payment_gateway"
	usecases "github.com/yescorihuela/deuna-payment-system/internal/application/usecases/payment_gateway"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/databases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api/handlers"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/requests"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/responses"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/repositories"
	http_client "github.com/yescorihuela/deuna-payment-system/internal/infrastructure/services/http"
	"github.com/yescorihuela/deuna-payment-system/internal/shared/utils"
)

func main() {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		panic(err)
	}

	db, err := databases.NewPostgresqlDbConnection(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pgTransactionRepository := repositories.NewPostgresqlTransactionRepository(db)

	pgRefundRepository := repositories.NewPostgresqlRefundRepository(db)
	httpClient := http_client.NewHttpClient[requests.PaymentRequest, responses.PaymentResponse](http_client.HttpClientSettings{
		Host:    fmt.Sprintf("%s:%s", config.HostAcquiringBank, config.HTTPServiceAcquiringBankPort),
		Timeout: config.TimeoutHTTPRequests,
	})

	paymentProcessUseCase := usecases.NewPaymentProcess(
		pgTransactionRepository,
		httpClient,
	)
	refundUseCase := usecases.NewRefundUseCase(pgRefundRepository)

	txHandler := handlers.NewTransactionHandler(paymentProcessUseCase)
	refundHandler := handlers.NewRefundHandler(refundUseCase)

	txApp := payment_gateway.NewApplication(
		txHandler,
		refundHandler,
		gin.Default(),
		config,
	)

	if err := txApp.Run(); err != nil {
		panic(err)
	}
}
