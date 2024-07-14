package application

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api"
)

type Application struct {
	txHandler     *api.TransactionHandler
	refundHandler *api.RefundHandler
	router        *gin.Engine
}

func NewApplication(
	txHandler *api.TransactionHandler,
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
	err := app.router.Run()
	return err
}
