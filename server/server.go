package server

import (
	"os"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
)

func Run() {
	app := echo.New()
	// app.Use(middleware.Secure())
	// app.Pre(middleware.HTTPSNonWWWRedirect())
	// app.Use(middleware.CORS())
	// app.Use(middleware.CSRF())
	// app.Use(middleware.Gzip())
	// app.Use(middleware.Timeout())
	// app.Use(middleware.BodyLimit("2M"))
	// app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	// app.Use(middleware.RequestID())
	// app.Use(middleware.Recover())
	// app.Use(middleware.Logger())
	app.Use(echoprometheus.NewMiddleware("yacco"))
	app.GET("/metrics", echoprometheus.NewHandler())

	app.Static("/", "./public/")
	app.Use(controller.Initialize())
	app.GET("/*", controller.Register(func(ctx *controller.Context) error { 
		return ctx.Html(view.Wildcard())
	}))
	storage.Connect(storage.Default())

	ServerRouters(app)
	app.Logger.Fatal(app.Start(os.Getenv("Port")))
}