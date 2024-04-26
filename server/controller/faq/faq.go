package faq

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	app.GET("/faq", controller.Register(index))
}