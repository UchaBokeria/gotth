package newser

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	newser := setting.Group("/newser")
	newser.POST("", controller.Register(NewserNew))
	newser.POST("/:id", controller.Register(Newser))
	newser.DELETE("/:id", controller.Register(NewserRemove))
}
