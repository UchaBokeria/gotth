package admin

import (
	"github.com/labstack/echo/v4"

	"main/server/controller/admin/category"
	"main/server/controller/admin/dashboard"
	"main/server/controller/admin/login"
	"main/server/controller/admin/product"
	"main/server/controller/admin/setting"
	"main/server/middleware"
)

func Register(app *echo.Echo) {
	admin := app.Group("admin")
	login.Register(admin)
	
	admin.Use(middleware.Auth())

	category.Register(admin)
	dashboard.Register(admin)
	product.Register(admin)
	setting.Register(admin)
}