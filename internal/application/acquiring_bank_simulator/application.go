package acquiring_bank_simulator

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api/handlers"
)

type Application struct {
	merchantHandler *handlers.AcquiringBankHandler
	router          *gin.Engine
}

func NewApplication(
	merchantHandler *handlers.AcquiringBankHandler,
	router *gin.Engine,
) *Application {
	return &Application{
		merchantHandler: merchantHandler,
		router:          router,
	}
}

func (app *Application) RegisterRoutes() {
	app.router.POST("/merchants/new", app.merchantHandler.New)
	app.router.PUT("/merchants/update/:merchant_code", app.merchantHandler.Update)
	app.router.PATCH("/merchants/change_status/:merchant_code", app.merchantHandler.ChangeStatus)
	app.router.GET("/merchants/by_code/:merchant_code", app.merchantHandler.GetByMerchantCode)
	app.router.GET("/merchants/by_id/:id", app.merchantHandler.GetById)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run() // add port from config
	return err
}
