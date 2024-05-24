package brancher

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	brancher := setting.Group("/brancher")
	brancher.POST("", controller.Register(BrancherNew))
	brancher.POST("/:id", controller.Register(Brancher))
	brancher.DELETE("/:id", controller.Register(BrancherRemove))
}