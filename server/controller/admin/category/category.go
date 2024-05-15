package category

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	Category := app.Group("/category")
	Category.GET("", controller.Register(index))
	// Category.POST("", controller.Register(create))
	// Category.PUT("", controller.Register(update))
	// Category.DELETE("", controller.Register(remove))
}