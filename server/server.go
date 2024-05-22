package server

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"main/server/common/controller"
	"main/server/common/globals"
	"main/server/common/storage"
)

func Run() {
	app := echo.New()
	

	app.Static("", "./public/")

	
    app.Pre(middleware.RemoveTrailingSlash())

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
	// app.Use(echoprometheus.NewMiddleware("yacco"))


	
	// app.GET("/metrics", echoprometheus.NewHandler())

	app.Use(controller.Initialize())
	storage.Connect(storage.Default())
	
	ServerRouters(app)
	data, _ := json.MarshalIndent(app.Routes(), "", "  ")
	os.WriteFile("routes.json", data, 0644)
	

	app.Logger.Fatal(app.Start(globals.Env.Port))
}