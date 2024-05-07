package middleware

import (
	"github.com/labstack/echo/v4"

	"main/build/view"
	"main/server/common/controller"
	"main/server/common/storage"
	"main/server/model"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Register(func(ctx *controller.Context) error {
			var Admin model.Users

			Email := ctx.Request().Header.Get("user")
			Token := ctx.Request().Header.Get("token")

			result := storage.DB.Last(&Admin, &model.Users{Email: Email, Token: Token})
			if result.Error == nil { return ctx.Renders(view.Login()) }

			ctx.Set("ISADMIN", true)
			ctx.Set("ADMIN", Admin)

			return next(ctx)
		})
	}
}

