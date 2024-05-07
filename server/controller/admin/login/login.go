package login

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Group) {
	app.GET("/login", controller.Register(index))
	app.POST("/login", controller.Register(login))
}