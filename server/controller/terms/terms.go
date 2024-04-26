package terms

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	app.GET("/terms", controller.Register(index))
}