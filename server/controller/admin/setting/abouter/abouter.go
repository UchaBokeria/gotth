package abouter

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	setting.POST("/termer", controller.Register(Termer))
	setting.POST("/abouter", controller.Register(Abouter))
}