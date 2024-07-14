package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application"
	"github.com/yescorihuela/deuna-payment-system/internal/application/usecases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/databases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/repositories"
)

func main() {
	db, err := databases.NewPostgresqlDbConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pgTransactionRepository := repositories.NewPostgresqlTransactionRepository(db)

	pgRefundRepository := repositories.NewPostgresqlRefundRepository(db)

	txUseCase := usecases.NewTransaction(pgTransactionRepository)
	refundUseCase := usecases.NewRefund(pgRefundRepository)

	txHandler := api.NewTransactionHandler(txUseCase)
	refundHandler := api.NewRefundHandler(refundUseCase)

	txApp := application.NewApplication(
		txHandler,
		refundHandler,
		gin.Default())

	if err := txApp.Run(); err != nil {
		panic(err)
	}
}
