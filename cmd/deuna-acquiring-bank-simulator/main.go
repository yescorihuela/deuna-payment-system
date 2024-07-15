package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/application/acquiring_bank_simulator"
	acquiring_bank_usecases "github.com/yescorihuela/deuna-payment-system/internal/application/usecases/acquiring_bank"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/databases"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api/handlers"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/repositories"
)

func main() {
	db, err := databases.NewPostgresqlDbConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pgMerchantRepository := repositories.NewPostgresqlMerchantRepository(db)

	merchantUseCase := acquiring_bank_usecases.NewMerchantUseCase(pgMerchantRepository)
	merchantHandler := handlers.NewAcquiringBankHandler(merchantUseCase)

	merchantApp := acquiring_bank_simulator.NewApplication(
		merchantHandler,
		gin.Default(),
	)

	if err := merchantApp.Run(); err != nil {
		panic(err)
	}
}
