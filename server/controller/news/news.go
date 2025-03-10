package news

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	app.GET("/news", controller.Register(index))
	app.GET("/news/:ID", controller.Register(detail))
}