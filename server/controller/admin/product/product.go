package product

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	Products := app.Group("/product")
	Products.GET("", controller.Register(index))
	Products.GET("/:id", controller.Register(indexByID))
	Products.POST("", controller.Register(ProductsNew))
	Products.PUT("", controller.Register(ProductsUpdate))
	Products.PATCH("/:id", controller.Register(ProductsStatus))
	Products.DELETE("/:id", controller.Register(ProductsRemove))
}