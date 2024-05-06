package products

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	app.GET("/products", controller.Register(index))
	app.GET("/products/list", controller.Register(list))
	app.GET("/products/:id", controller.Register(detail))
}