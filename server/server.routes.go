package server

import (
	"github.com/labstack/echo/v4"

	"main/server/controller/about"
	"main/server/controller/admin"
	"main/server/controller/branches"
	"main/server/controller/categories"
	"main/server/controller/faq"
	"main/server/controller/landing"
	"main/server/controller/news"
	"main/server/controller/products"
	"main/server/controller/terms"
	"main/server/middleware"
)

func ServerRouters(app *echo.Echo) {
	admin.Register(app)

	app.Use(middleware.Interface())
	landing.Register(app)
	categories.Register(app)
	products.Register(app)
	branches.Register(app)
	news.Register(app)
	faq.Register(app)
	about.Register(app)
	terms.Register(app)
}