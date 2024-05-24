package faqers

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(setting *echo.Group) {
	faqers := setting.Group("/faqers")
	faqers.POST("", controller.Register(FaqersNew))
	faqers.POST("/:id", controller.Register(Faqers))
	faqers.DELETE("/:id", controller.Register(FaqersRemove))
}