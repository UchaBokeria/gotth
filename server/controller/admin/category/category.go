package category

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	app.GET("", controller.Register(index))
	Category := app.Group("/category")
	Category.GET("", controller.Register(index))
	Category.GET("/:id", controller.Register(indexByID))
	Category.POST("", controller.Register(CategoryNew))
	Category.PUT("", controller.Register(CategoryUpdate))
	Category.PATCH("/:id", controller.Register(CategoryStatus))
	Category.DELETE("/:id", controller.Register(CategoryRemove))
}