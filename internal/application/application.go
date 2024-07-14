package application

import (
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/deuna-payment-system/internal/infrastructure/http/api"
)

type Application struct {
	txHandler *api.TransactionHandler
	router    *gin.Engine
}

func NewApplication(
	txHandler *api.TransactionHandler,
	router *gin.Engine,
) *Application {
	return &Application{
		txHandler: txHandler,
		router:    router,
	}
}

func (app *Application) RegisterRoutes() {
	app.router.POST("/transaction/new", app.txHandler.Create)
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run()
	return err
}
