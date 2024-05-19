package upload

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	app.POST("/upload", controller.Register(FileUpload))
}