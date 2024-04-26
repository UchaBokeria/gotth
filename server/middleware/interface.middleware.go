package middleware

import (
	"github.com/labstack/echo/v4"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func Interface() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Register(func(ctx *controller.Context) error {
			if ctx.IsHtmx() { return next(ctx) }

			var Interface model.Interface
			result := storage.DB.Preload("Contact").Preload("SocialMedia").Last(&Interface)
			if result.Error != nil { return ctx.Html(view.ErrorPage()) }

			ctx.Set("Interface", Interface)
			return next(ctx)
		})
	}
}

