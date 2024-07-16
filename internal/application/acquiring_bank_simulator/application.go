package acquiring_bank_simulator

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api/handlers"
	"github.com/yescorihuela/deuna-payment-system/internal/shared/utils"
)

type Application struct {
	merchantHandler *handlers.AcquiringBankHandler
	router          *gin.Engine
	config          utils.Config
}

func NewApplication(
	merchantHandler *handlers.AcquiringBankHandler,
	router *gin.Engine,
	config utils.Config,
) *Application {
	return &Application{
		merchantHandler: merchantHandler,
		router:          router,
		config:          config,
	}
}

func (app *Application) RegisterRoutes() {
	v1 := app.router.Group("/v1")
	v1.POST("/merchants/new", app.merchantHandler.New)
	v1.PUT("/merchants/update/:merchant_code", app.merchantHandler.Update)
	v1.PATCH("/merchants/change_status/:merchant_code", app.merchantHandler.ChangeStatus)
	v1.GET("/merchants/by_code/:merchant_code", app.merchantHandler.GetByMerchantCode)
	v1.GET("/merchants/by_id/:id", app.merchantHandler.GetById)

	app.router.POST("/transaction", app.merchantHandler.ExecuteTransaction)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run(fmt.Sprintf(":%s", app.config.HTTPServiceAcquiringBank))
	return err
}
