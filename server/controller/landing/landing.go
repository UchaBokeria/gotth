package landing

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	app.GET("/", controller.Register(index))

	router := app.Group("/landing")
	router.GET("/", controller.Register(index))
}