package reasoner

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	reasoner := setting.Group("/reasoner")
	reasoner.POST("", controller.Register(ReasonerNew))
	reasoner.POST("/:id", controller.Register(Reasoner))
	reasoner.DELETE("/:id", controller.Register(ReasonerRemove))
}
