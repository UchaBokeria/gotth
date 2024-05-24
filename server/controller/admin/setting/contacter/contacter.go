package contacter

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	setting.POST("/socials", controller.Register(Socials))
	setting.POST("/contacts", controller.Register(Contact))
}