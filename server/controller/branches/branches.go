package branches

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	router := app.Group("branches")
	router.GET("/", controller.Register(index))
}