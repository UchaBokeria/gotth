package chat

import (
	"github.com/labstack/echo/v4"

	"main/server/common/controller"
)

func Register(app *echo.Echo) {
	Chat := app.Group("/chat")
	Chat.POST("", controller.Register(index))
	Chat.POST("/mailStrategy", controller.Register(MailStrategy))
	// Chat.GET("/test", controller.Register(func(ctx *controller.Context) error {
	// 	return ctx.Html(view.ChatTest())
	// }))
	// Chat.GET("/new", controller.Register(NewChat))
	// Chat.GET("/:id", controller.Register(OldChat))
	// Chat.GET(":id", controller.Register(index))
}