package setting

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	app.POST("/setting/slideshower", controller.Register(SlideshowerNew))
	app.POST("/setting/slideshower/:id", controller.Register(Slideshower))
	app.DELETE("/setting/slideshower/:id", controller.Register(SlideRemove))
}