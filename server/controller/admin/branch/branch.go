package branch

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	app.GET("/branch", controller.Register(index))
}