package payment_gateway

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api/handlers"
	"github.com/yescorihuela/deuna-payment-system/internal/shared/utils"
)

type Application struct {
	txHandler     *handlers.PaymentHandler
	refundHandler *handlers.RefundHandler
	router        *gin.Engine
	config        utils.Config
}

func NewApplication(
	txHandler *handlers.PaymentHandler,
	refundHandler *handlers.RefundHandler,
	router *gin.Engine,
	config utils.Config,
) *Application {
	return &Application{
		txHandler:     txHandler,
		refundHandler: refundHandler,
		router:        router,
		config:        config,
	}
}

func (app *Application) RegisterRoutes() {
	v1 := app.router.Group("/v1")
	v1.POST("/payments/new", app.txHandler.Create)
	v1.POST("/payments/refund", app.refundHandler.Create)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run(fmt.Sprintf(":%s", app.config.HTTPServicePaymentPort))
	return err
}
