package slideshower

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	slideshower := setting.Group("/slideshower")
	slideshower.POST("", controller.Register(SlideshowerNew))
	slideshower.POST("/:id", controller.Register(Slideshower))
	slideshower.DELETE("/:id", controller.Register(SlideRemove))
}
