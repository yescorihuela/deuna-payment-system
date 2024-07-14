package payment_gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api"
)

type Application struct {
	txHandler     *api.PaymentHandler
	refundHandler *api.RefundHandler
	router        *gin.Engine
}

func NewApplication(
	txHandler *api.PaymentHandler,
	refundHandler *api.RefundHandler,
	router *gin.Engine,
) *Application {
	return &Application{
		txHandler:     txHandler,
		refundHandler: refundHandler,
		router:        router,
	}
}

func (app *Application) RegisterRoutes() {
	app.router.POST("/payments/new", app.txHandler.Create)
	app.router.POST("/payments/refund", app.refundHandler.Create)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run() // add port from config
	return err
}
