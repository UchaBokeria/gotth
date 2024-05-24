package setting

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
	"main/server/controller/admin/setting/abouter"
	"main/server/controller/admin/setting/brancher"
	"main/server/controller/admin/setting/contacter"
	"main/server/controller/admin/setting/faqers"
	"main/server/controller/admin/setting/newser"
	"main/server/controller/admin/setting/reasoner"
	"main/server/controller/admin/setting/slideshower"
)

func Register(admin *echo.Group) {
	admin.GET("/settings/", controller.Register(index))
	admin.GET("/settings/:tab", controller.Register(index))

	setting := admin.Group("/setting")

	abouter.Register(setting)
	brancher.Register(setting)
	contacter.Register(setting)
	faqers.Register(setting)
	newser.Register(setting)
	reasoner.Register(setting)
	slideshower.Register(setting)
}