package acquiring_bank_simulator

import "github.com/gin-gonic/gin"

type Application struct {
	router *gin.Engine
}

func NewApplication(
	router *gin.Engine,
) *Application {
	return &Application{
		router: router,
	}
}

func (app *Application) RegisterRoutes() {
	app.router.POST("/merchants/new")
	app.router.PUT("/merchants/update")
	app.router.PATCH("/merchants/change_status/:merchant_code")
	app.router.GET("/merchants/by_code/:merchant_code")
	app.router.GET("/merchants/by_id/:id")
}

func (app *Application) Bootstrapping() {
	app.RegisterRoutes()
}

func (app *Application) Run() error {
	app.Bootstrapping()
	err := app.router.Run() // add port from config
	return err
}
