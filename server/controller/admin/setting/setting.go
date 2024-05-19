package setting

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	app.GET("/setting", controller.Register(index))
	app.POST("/setting/socials", controller.Register(Socials))
	app.POST("/setting/contacts", controller.Register(Contact))

	app.POST("/setting/slideshower", controller.Register(SlideshowerNew))
	app.POST("/setting/slideshower/:id", controller.Register(Slideshower))
	app.DELETE("/setting/slideshower/:id", controller.Register(SlideRemove))

	app.POST("/setting/reasoner", controller.Register(ReasonerNew))
	app.POST("/setting/reasoner/:id", controller.Register(Reasoner))
	app.DELETE("/setting/reasoner/:id", controller.Register(ReasonerRemove))

	app.POST("/setting/faqers", controller.Register(FaqersNew))
	app.POST("/setting/faqers/:id", controller.Register(Faqers))
	app.DELETE("/setting/faqers/:id", controller.Register(FaqersRemove))

	app.POST("/setting/newser", controller.Register(NewserNew))
	app.POST("/setting/newser/:id", controller.Register(Newser))
	app.DELETE("/setting/newser/:id", controller.Register(NewserRemove))
}